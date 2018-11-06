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
import torch
from pathlib import Path
from io import BytesIO
import sys
import uvicorn
import aiohttp
import asyncio
import base64
from PIL import Image


async def get_bytes(url):
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            return await response.read()


app = Starlette()
app.add_middleware(CORSMiddleware, allow_origins=['*'], allow_headers=['*'])

path = Path("./")
classes = ['Brazil', 'Armenia', 'Thailand', 'Algeria', 'China', 'United+Kingdom', 'South+Korea', 'Hungary', 'Central+African+Republic', 'Poland', 'Venezuela', 'Germany', 'France', 'Argentina', 'Russia', 'Cuba', 'Spain', 'Mexico', 'Japan', 'Philippines', 'Ukraine', 'Italy', 'Macedonia', 'India', 'Portugal', 'Ghana', 'Madagascar', 'United+States', 'Indonesia', 'Slovakia', 'Malaysia', 'Czech+Republic', 'Romania', 'Sri+Lanka', 'Bulgaria', 'Guinea', 'Colombia', 'Vietnam', 'Turkey', 'Seychelles', 'Pakistan', 'Chad', 'Cambodia', 'Iraq', 'Finland', 'Belarus', 'Afghanistan', 'Burkina+Faso', 'South+Africa', 'Australia', 'Peru', 'Kenya', 'Tunisia', 'Denmark', 'Canada', 'Namibia', 'Nigeria', 'Zambia', 'Iran', 'El+Salvador', 'Myanmar', 'Greece', 'Ethiopia', 'Sweden', 'Oman', 'Egypt', 'Bhutan', 'Belgium', 'Azerbaijan', 'Mozambique', 'Libya', 'Ivory+Coast', 'Croatia', 'East+Timor', 'Niger', 'Netherlands', 'Albania', 'Dominican+Republic', 'Somalia', 'Tanzania', 'Kosovo', 'French+Guiana', 'Ireland', 'Cameroon', 'Norway', 'French+Polynesia', 'Ecuador', 'Botswana', 'Angola', 'Syria', 'Costa+Rica', 'Morocco', 'Bangladesh', 'Yemen', 'Serbia', 'Lithuania', 'Uganda', 'Taiwan', 'Bosnia+and+Herzegovina', 'North+Korea', 'Israel', 'Benin', 'Slovenia', 'Papua+New+Guinea', 'Rwanda', 'Zimbabwe', 'Martinique', 'Moldova', 'Senegal', 'Gambia', 'Kuwait', 'Sudan', 'Malawi', 'Paraguay', 'Austria', 'Georgia', 'Wallis+and+Futuna', 'Kazakhstan', 'Panama', 'Mongolia', 'New+Zealand', 'Tajikistan', 'Switzerland', 'Sierra+Leone', 'Uzbekistan', 'Guam', 'Mauritius', 'Uruguay', 'Mayotte', 'Democratic+Republic+of+the+Congo', 'Guatemala', 'Jordan', 'Bolivia', 'Puerto+Rico', 'Honduras', 'Chile', 'Republic+of+the+Congo', 'Nicaragua', 'Guyana', 'Malta', 'Liberia', 'Mali', 'Trinidad+and+Tobago', 'Togo', 'Qatar', '"Bonaire,+Saint+Eustatius+and+Saba"', 'Lebanon', 'Faroe+Islands', 'Maldives', 'United+Arab+Emirates', 'Gabon', 'Laos', 'Estonia', 'Turkmenistan', 'Mauritania', 'Iceland', 'Saudi+Arabia', 'Hong+Kong', 'Haiti', 'Greenland', 'Brunei', 'Nepal', 'Bahrain', 'Cayman+Islands', 'Cape+Verde', 'Fiji', 'Suriname', 'Aruba', 'Pitcairn', 'Guinea-Bissau', 'Cook+Islands', 'Tonga', 'Eritrea', 'Niue', 'Cyprus', 'Palestinian+Territory', 'Gibraltar', 'Kyrgyzstan', 'Djibouti', 'Burundi', 'Saint+Helena', 'Tuvalu', 'Montenegro', 'Guadeloupe', 'South+Sudan', 'Falkland+Islands', 'Jamaica', 'Latvia', 'New+Caledonia', 'Samoa', 'Reunion', 'Curacao', 'San+Marino', 'Comoros', 'Guernsey', 'Lesotho', 'Sao+Tome+and+Principe', 'Northern+Mariana+Islands', 'Belize', 'Swaziland', 'Equatorial+Guinea', 'Palau', 'Barbados', 'Isle+of+Man', 'Cocos+Islands', 'Aland+Islands', 'Saint+Vincent+and+the+Grenadines', 'Luxembourg', 'Western+Sahara', 'Anguilla', 'Liechtenstein', 'Macao', 'Saint+Pierre+and+Miquelon', 'Andorra', 'French+Southern+Territories', 'British+Virgin+Islands', 'Montserrat', 'Nauru', 'Saint+Barthelemy', 'Bahamas', 'South+Georgia+and+the+South+Sandwich+Islands', 'Dominica', 'Singapore', 'Vanuatu', 'Bermuda', 'Jersey', 'Christmas+Island', 'Grenada', '+D.C.",United+States', 'Solomon+Islands', 'Monaco', 'Kiribati', 'Norfolk+Island', 'Vatican', 'Saint+Lucia', 'Saint+Kitts+and+Nevis', 'Turks+and+Caicos+Islands', 'Micronesia', 'U.S.+Virgin+Islands', 'Antigua+and+Barbuda', 'American+Samoa', 'Saint+Martin', 'Svalbard+and+Jan+Mayen', 'Sint+Maarten']

fake_data = ImageDataBunch.single_from_classes(path, classes, tfms=get_transforms(), size=224).normalize(imagenet_stats)
learn = create_cnn(fake_data, models.resnet34)
learn.model.load_state_dict(torch.load('resnet50-big-finetuned-bs64.pth', map_location='cpu'))



@app.route("/upload", methods=["POST"])
async def upload(request):
    data = await request.form()
    s = data["file"]
    return predict_image_from_string(s)


def predict_image_from_string(s):
    b = base64.b64decode(s)
    img = open_image(BytesIO(b))
    pred_class, pred_idx, outputs = learn.predict(img)
    return JSONResponse({
        "predictions": pred_class.replace('+', ' ').strip()
    })





if __name__ == "__main__":
    if "serve" in sys.argv:
        uvicorn.run(app, host="0.0.0.0", port=8008)
