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
classes = ['+Zimbabwe', '+China', '+Mexico', '+Nigeria', '+Israel', '+Benin', '+Australia', '+Japan', '+India', '+Saudi+Arabia', '+Philippines', '+Iran', '+New+Zealand', '+United+States', '+Brazil', '+Singapore', '+Morocco', '+Russia', '+Cameroon', '+Indonesia', '+Peru', '+Kenya', '+Algeria', '+Pakistan', '+Canada', '+Syria', '+Spain', '+France', '+United+Kingdom', '+Ukraine', '+Malaysia', '+Poland', '+Uzbekistan', '+Venezuela', '+South+Africa', '+Colombia', '+Egypt', '+Mali', '+South+Korea', '+Zambia', '+Germany', '+Bolivia', '+Argentina', '+Thailand', '+Switzerland', '+Azerbaijan', '+Turkey', '+Burkina+Faso', '+Italy', '+Dominican+Republic', '+Malawi', '+Laos', '+Belarus', '+Armenia', '+Dem+Rep+of+Congo', '+Greece', '+Kazakhstan', '+North+Korea', '+Nicaragua', '+Bangladesh', '+Netherlands', '+China+-+Hong+Kong', '+Ghana', '+Madagascar', '+Viet+Nam', '+Chile', '+Jamaica', '+Tanzania', '+Kyrgyzstan', '+Croatia', '+Georgia', '+Iraq', '+Ecuador', '+Guinea', '+Hungary', '+Bulgaria', '+Myanmar', '+Serbia', '+Cambodia', '+Romania', '+Libya', '+Haiti', '+Latvia', '+United+Arab+Emirates', '+Sudan', '+Senegal', '+Sweden', '+Finland', '+Cuba', '+Ethiopia', '+Honduras', '+Ireland', '+Yemen', '+Sri+Lanka', '+Lithuania', '+Moldova', '+Portugal', '+Belgium', '+Guatemala', '+Chad', '+Paraguay', '+Jordan', '+Tunisia', '+Uganda', '+Somalia', '+Mozambique', '+Mongolia', '+Afghanistan', '+Niger', '+Czech+Republic', '+Mauritania', '+Nepal', '+Angola', '+Norway', '+Uruguay', '+Austria', "+Côte+d'Ivoire", '+Taiwan+(China+ROC)']

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
