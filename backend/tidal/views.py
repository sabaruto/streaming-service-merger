from django.http import HttpResponse
from django.views import View

import base64
import json
import requests
from requests.auth import HTTPBasicAuth


client_id = "W4kUNzVf3QbDiAjR"
client_secrets = "PgbZAT3kR4kBIDwTnFu5xu5hOWLc3Gy9KVDXQTUwVlY="


def index(request):
    return HttpResponse("Hello World")


class AuthView(View):
    def get(self, request, *args, **kwargs):
        client_id_encoded = base64.encodebytes(bytes(client_id, "utf-8"))
        client_secret_encoded = base64.encodebytes(bytes(client_secrets, "utf-8"))


        headers = {"Content-Type": "application/x-www-form-urlencoded"}
        response = requests.post(
            "https://auth.tidal.com/v1/oauth2/token",
            headers=headers,
            auth=HTTPBasicAuth(client_id, client_secrets),
            data="grant_type=client_credentials",
            timeout=10,
        )

        access_token = json.loads(response.content)

        return HttpResponse(access_token["accesss_token"])
