from starlette.applications import Starlette
from starlette.responses import JSONResponse, HTMLResponse, RedirectResponse
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


async def get_bytes(url):
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            return await response.read()


app = Starlette()

path = Path("./")
classes = ['+Indonesia', '+Argentina', '+China', '+Canada', '+United+States', '+Japan', '+Syria', '+Malaysia', '+Brazil', '+Myanmar', '+Turkey', '+India', '+Ukraine', '+Iran', '+Mexico', '+Egypt', '+Saudi+Arabia', '+Ghana', '+Belarus', '+China+-+Hong+Kong', '+South+Korea', '+Colombia', '+Morocco', '+Philippines', '+South+Africa', '+Germany', '+United+Kingdom', '+Italy', '+Peru', '+Nepal', '+Sudan', '+Chile', '+Venezuela', '+Pakistan', '+Jordan', '+Sweden', '+Serbia', '+Thailand', '+Spain', '+Mali', '+Niger', '+Libya', '+Russia', '+Zambia', '+Nigeria', '+Guatemala', '+Georgia', '+Laos', '+Viet+Nam', '+Bangladesh', '+Czech+Republic', '+Belgium', '+Norway', '+Bulgaria', '+North+Korea', '+Angola', '+Australia', '+Romania', '+Algeria', '+Kenya', '+Cameroon', '+Poland', '+Uruguay', '+Dominican+Republic', '+Netherlands', "+Côte+d'Ivoire", '+Senegal', '+Jamaica', '+Somalia', '+United+Arab+Emirates', '+Ireland', '+Bolivia', '+Benin', '+Taiwan+(China+ROC)', '+Ecuador', '+Azerbaijan', '+Honduras', '+Tanzania', '+Zimbabwe', '+Cuba', '+Nicaragua', '+Kazakhstan', '+France', '+Madagascar', '+Croatia', '+Haiti', '+Mozambique', '+Moldova', '+Hungary', '+Latvia', '+Greece', '+Lithuania', '+Uganda', '+Cambodia', '+Sri+Lanka', '+Mauritania', '+Finland', '+Iraq', '+Mongolia', '+New+Zealand', '+Afghanistan', '+Israel', '+Singapore', '+Dem+Rep+of+Congo', '+Austria', '+Paraguay', '+Portugal', '+Guinea', '+Malawi', '+Yemen', '+Chad', '+Kyrgyzstan', '+Armenia', '+Burkina+Faso', '+Tunisia', '+Uzbekistan', '+Switzerland', '+Ethiopia']

fake_data = ImageDataBunch.single_from_classes(path, classes, tfms=get_transforms(), size=224).normalize(imagenet_stats)
learn = create_cnn(fake_data, models.resnet34)
learn.model.load_state_dict(torch.load('finetuned_cities.pth', map_location='cpu'))



@app.route("/upload", methods=["POST"])
async def upload(request):
    data = await request.form()
    bytes = await (data["file"].read())
    return predict_image_from_bytes(bytes)


@app.route("/classify-url", methods=["GET"])
async def classify_url(request):
    bytes = await get_bytes(request.query_params["url"])
    return predict_image_from_bytes(bytes)


def predict_image_from_bytes(bytes):
    img = open_image(BytesIO(bytes))
    pred_class,pred_idx,outputs = learn.predict(img)
    return JSONResponse({
        "predictions": pred_class
    })


@app.route("/")
def form(request):
    return HTMLResponse(
        """
        <form action="/upload" method="post" enctype="multipart/form-data">
            Select image to upload:
            <input type="file" name="file">
            <input type="submit" value="Upload Image">
        </form>
        Or submit a URL:
        <form action="/classify-url" method="get">
            <input type="url" name="url">
            <input type="submit" value="Fetch and analyze image">
        </form>
    """)


@app.route("/form")
def redirect_to_homepage(request):
    return RedirectResponse("/")


if __name__ == "__main__":
    if "serve" in sys.argv:
        uvicorn.run(app, host="0.0.0.0", port=8008)