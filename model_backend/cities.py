from starlette.applications import Starlette
from starlette.responses import JSONResponse, HTMLResponse, RedirectResponse
from starlette.middleware.cors import CORSMiddleware
from fastai.vision import (
    ImageDataBunch,
    create_cnn,
    open_image,
    get_transforms,
    models,
    imagenet_stats
)
from fastai import (
    Hook,
    hook_output
)
import torch
from pathlib import Path
from io import BytesIO
import sys
import uvicorn
import aiohttp
import asyncio
import base64
from PIL import Image
import matplotlib.pyplot as plt
import scipy.ndimage
import numpy as np


async def get_bytes(url):
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            return await response.read()


app = Starlette()
app.add_middleware(CORSMiddleware, allow_origins=['*'], allow_headers=['*'])

path = Path("./")
classes = ['Brazil', 'Armenia', 'Thailand', 'Algeria', 'China', 'United+Kingdom', 'South+Korea', 'Hungary', 'Central+African+Republic', 'Poland', 'Venezuela', 'Germany', 'France', 'Argentina', 'Russia', 'Cuba', 'Spain', 'Mexico', 'Japan', 'Philippines', 'Ukraine', 'Italy', 'Macedonia', 'India', 'Portugal', 'Ghana', 'Madagascar', 'United+States', 'Indonesia', 'Slovakia', 'Malaysia', 'Czech+Republic', 'Romania', 'Sri+Lanka', 'Bulgaria', 'Guinea', 'Colombia', 'Vietnam', 'Turkey', 'Seychelles', 'Pakistan', 'Chad', 'Cambodia', 'Iraq', 'Finland', 'Belarus', 'Afghanistan', 'Burkina+Faso', 'South+Africa', 'Australia', 'Peru', 'Kenya', 'Tunisia', 'Denmark', 'Canada', 'Namibia', 'Nigeria', 'Zambia', 'Iran', 'El+Salvador', 'Myanmar', 'Greece', 'Ethiopia', 'Sweden', 'Oman', 'Egypt', 'Bhutan', 'Belgium', 'Azerbaijan', 'Mozambique', 'Libya', 'Ivory+Coast', 'Croatia', 'East+Timor', 'Niger', 'Netherlands', 'Albania', 'Dominican+Republic', 'Somalia', 'Tanzania', 'Kosovo', 'French+Guiana', 'Ireland', 'Cameroon', 'Norway', 'French+Polynesia', 'Ecuador', 'Botswana', 'Angola', 'Syria', 'Costa+Rica', 'Morocco', 'Bangladesh', 'Yemen', 'Serbia', 'Lithuania', 'Uganda', 'Taiwan', 'Bosnia+and+Herzegovina', 'North+Korea', 'Israel', 'Benin', 'Slovenia', 'Papua+New+Guinea', 'Rwanda', 'Zimbabwe', 'Martinique', 'Moldova', 'Senegal', 'Gambia', 'Kuwait', 'Sudan', 'Malawi', 'Paraguay', 'Austria', 'Georgia', 'Wallis+and+Futuna', 'Kazakhstan', 'Panama', 'Mongolia', 'New+Zealand', 'Tajikistan', 'Switzerland', 'Sierra+Leone', 'Uzbekistan', 'Guam', 'Mauritius', 'Uruguay', 'Mayotte', 'Democratic+Republic+of+the+Congo', 'Guatemala', 'Jordan', 'Bolivia', 'Puerto+Rico', 'Honduras', 'Chile', 'Republic+of+the+Congo', 'Nicaragua', 'Guyana', 'Malta', 'Liberia', 'Mali', 'Trinidad+and+Tobago', 'Togo', 'Qatar', '"Bonaire,+Saint+Eustatius+and+Saba"', 'Lebanon', 'Faroe+Islands', 'Maldives', 'United+Arab+Emirates', 'Gabon', 'Laos', 'Estonia', 'Turkmenistan', 'Mauritania', 'Iceland', 'Saudi+Arabia', 'Hong+Kong', 'Haiti', 'Greenland', 'Brunei', 'Nepal', 'Bahrain', 'Cayman+Islands', 'Cape+Verde', 'Fiji', 'Suriname', 'Aruba', 'Pitcairn', 'Guinea-Bissau', 'Cook+Islands', 'Tonga', 'Eritrea', 'Niue', 'Cyprus', 'Palestinian+Territory', 'Gibraltar', 'Kyrgyzstan', 'Djibouti', 'Burundi', 'Saint+Helena', 'Tuvalu', 'Montenegro', 'Guadeloupe', 'South+Sudan', 'Falkland+Islands', 'Jamaica', 'Latvia', 'New+Caledonia', 'Samoa', 'Reunion', 'Curacao', 'San+Marino', 'Comoros', 'Guernsey', 'Lesotho', 'Sao+Tome+and+Principe', 'Northern+Mariana+Islands', 'Belize', 'Swaziland', 'Equatorial+Guinea', 'Palau', 'Barbados', 'Isle+of+Man', 'Cocos+Islands', 'Aland+Islands', 'Saint+Vincent+and+the+Grenadines', 'Luxembourg', 'Western+Sahara', 'Anguilla', 'Liechtenstein', 'Macao', 'Saint+Pierre+and+Miquelon', 'Andorra', 'French+Southern+Territories', 'British+Virgin+Islands', 'Montserrat', 'Nauru', 'Saint+Barthelemy', 'Bahamas', 'South+Georgia+and+the+South+Sandwich+Islands', 'Dominica', 'Singapore', 'Vanuatu', 'Bermuda', 'Jersey', 'Christmas+Island', 'Grenada', '+D.C.",United+States', 'Solomon+Islands', 'Monaco', 'Kiribati', 'Norfolk+Island', 'Vatican', 'Saint+Lucia', 'Saint+Kitts+and+Nevis', 'Turks+and+Caicos+Islands', 'Micronesia', 'U.S.+Virgin+Islands', 'Antigua+and+Barbuda', 'American+Samoa', 'Saint+Martin', 'Svalbard+and+Jan+Mayen', 'Sint+Maarten']

fake_data = ImageDataBunch.single_from_classes(path, classes, tfms=get_transforms(), size=224).normalize(imagenet_stats)
learn = create_cnn(fake_data, models.resnet50)
learn.model.load_state_dict(torch.load('resnet50-big-finetuned-bs64.pth', map_location='cpu'))



@app.route("/upload", methods=["POST"])
async def upload(request):
    data = await request.form()
    s = data["file"]
    return predict_image_from_string(s)

@app.route("/heatmap", methods=["POST"])
async def heatmap(request):
    data = await request.form()
    s = data["file"]
    return heatmap_from_string(s)

def predict_image_from_string(s):
    b = base64.b64decode(s)
    img = open_image(BytesIO(b))
    pred_class, pred_idx, outputs = learn.predict(img)
    return JSONResponse({
        "predictions": pred_class.replace('+', ' ').strip()
    })

def heatmap_from_string(s):
    b = base64.b64decode(s)
    img = open_image(BytesIO(b))
    pred_class, pred_idx, outputs = learn.predict(img)
    img = img.px.reshape(1, 3, 224, 224)

    upsampled = run_gradcam(img)

    figdata_png = upsampled_to_b64bytes(upsampled, img)

    return JSONResponse({
        "predictions": figdata_png.decode('ascii')
    })

def upsampled_to_b64bytes(upsampled, img):
    """
    this combines upsampled heatmap and img
    and returns b64 encoded bytes for the image
    """
    figfile = BytesIO()

    fig = plt.figure(frameon=False)
    fig.set_size_inches(2,2)

    # all this to remove borders
    ax = plt.Axes(fig, [0., 0., 1., 1.])
    ax.set_axis_off()
    fig.add_axes(ax)
    ax.imshow(image_from_tensor(img))
    ax.imshow(upsampled, alpha=.6)
    plt.savefig(figfile, format='png', aspect='normal')

    figfile.seek(0) 
    figdata_png = base64.b64encode(figfile.getvalue())

    return figdata_png


def run_gradcam(img):
    """
    returns the heatmap for the given image
    """
    # last bottleneck module
    target_layer = learn.model[0][7][2]

    fmap_hook, gradient_hook = create_hooks(target_layer)
    run_backprop_once(img)

    gradient  = next(iter(gradient_hook.stored))
    linearization = gradient.cpu().numpy().sum((2, 3)).reshape(-1)
    fmaps = fmap_hook.stored.cpu().numpy()
    fmaps = fmaps.reshape(2048, 7, 7)

    hm = np.maximum(0, np.einsum('i, ijk',linearization, fmaps))
    upsampled = scipy.ndimage.zoom(hm, 32)

    return upsampled


def image_from_tensor(imagetensor):
    numpied = torch.squeeze(imagetensor)
    numpied = np.moveaxis(numpied.cpu().numpy(), 0 , -1)
    numpied = numpied - np.min(numpied)
    numpied = numpied/np.max(numpied)
    return numpied

def create_hooks(target_layer):
    feature_maps = hook_output(target_layer)
    gradient_hook = Hook(target_layer, gradient_torch_hook, is_forward=False)

    return feature_maps, gradient_hook 


def run_backprop_once(img):
    # forward
    out = learn.model(img)

    # gradient wrt the predicted class only
    onehot = torch.zeros(learn.data.c)
    torch.argmax(out)
    onehot[torch.argmax(out)] = 1.0

    # backwrd
    out.backward(gradient=onehot.reshape(1, -1))


def gradient_torch_hook(self, grad_input, grad_output):
    return grad_input


if __name__ == "__main__":
    if "serve" in sys.argv:
        uvicorn.run(app, host="0.0.0.0", port=8008)
