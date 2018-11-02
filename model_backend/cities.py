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
classes = ['+Indonesia', '+Argentina', '+China', '+Canada', '+United+States', '+Japan', '+Syria', '+Malaysia', '+Brazil', '+Myanmar', '+Turkey', '+India', '+Ukraine', '+Iran', '+Mexico', '+Egypt', '+Saudi+Arabia', '+Ghana', '+Belarus', '+China+-+Hong+Kong', '+South+Korea', '+Colombia', '+Morocco', '+Philippines', '+South+Africa', '+Germany', '+United+Kingdom', '+Italy', '+Peru', '+Nepal', '+Sudan', '+Chile', '+Venezuela', '+Pakistan', '+Jordan', '+Sweden', '+Serbia', '+Thailand', '+Spain', '+Mali', '+Niger', '+Libya', '+Russia', '+Zambia', '+Nigeria', '+Guatemala', '+Georgia', '+Laos', '+Viet+Nam', '+Bangladesh', '+Czech+Republic', '+Belgium', '+Norway', '+Bulgaria', '+North+Korea', '+Angola', '+Australia', '+Romania', '+Algeria', '+Kenya', '+Cameroon', '+Poland', '+Uruguay', '+Dominican+Republic', '+Netherlands', "+Côte+d'Ivoire", '+Senegal', '+Jamaica', '+Somalia', '+United+Arab+Emirates', '+Ireland', '+Bolivia', '+Benin', '+Taiwan+(China+ROC)', '+Ecuador', '+Azerbaijan', '+Honduras', '+Tanzania', '+Zimbabwe', '+Cuba', '+Nicaragua', '+Kazakhstan', '+France', '+Madagascar', '+Croatia', '+Haiti', '+Mozambique', '+Moldova', '+Hungary', '+Latvia', '+Greece', '+Lithuania', '+Uganda', '+Cambodia', '+Sri+Lanka', '+Mauritania', '+Finland', '+Iraq', '+Mongolia', '+New+Zealand', '+Afghanistan', '+Israel', '+Singapore', '+Dem+Rep+of+Congo', '+Austria', '+Paraguay', '+Portugal', '+Guinea', '+Malawi', '+Yemen', '+Chad', '+Kyrgyzstan', '+Armenia', '+Burkina+Faso', '+Tunisia', '+Uzbekistan', '+Switzerland', '+Ethiopia']

fake_data = ImageDataBunch.single_from_classes(path, classes, tfms=get_transforms(), size=224).normalize(imagenet_stats)
learn = create_cnn(fake_data, models.resnet34)
learn.model.load_state_dict(torch.load('finetuned_cities.pth', map_location='cpu'))



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