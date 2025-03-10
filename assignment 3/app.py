from flask import Flask

app = Flask(__name__)

@app.route("/")
def helloFunc():
    print("Страница открыта")
    return "<h1>Hello, world!!!</h1>"


if __name__ == "__main__":
    print("Запуск сервера...")
    app.run(host="0.0.0.0", port=8080, debug=True)