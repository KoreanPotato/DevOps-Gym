from fastapi import FastAPI, Request
from fastapi.responses import PlainTextResponse
import json

app = FastAPI()

@app.get("/api/hello", response_class=PlainTextResponse)
async def get_hello(request: Request):
    body = await request.body()
    
    if not body:
        return "Name is required"

    data = json.loads(body)
    name = data.get("name")
    
    return f"Hello, {name}!"
