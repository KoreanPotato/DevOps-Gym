from fastapi import FastAPI, Request
import json

app = FastAPI()

@app.get("/api/hello")
async def getHello(request: Request):
    body = await request.body()
    
    if not body:  
        return {"error": "name is required"}
    
    data = json.loads(body) 
    name = data.get("name") 
    
    return {"message": f"Hello, {name}!"}
