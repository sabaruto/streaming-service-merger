from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.models import User
from django.http import HttpRequest, HttpResponse, QueryDict
from django.views import View


def check_login_info(dict: QueryDict) -> tuple[str, str] | HttpResponse:
    response = HttpResponse()

    try:
        username = dict["username"]
        password = dict["password"]
    except KeyError:
        response.content = "Username and / or password not given."
        response.status_code = 400
        return response

    return username, password


def authenticate_request(request: HttpRequest) -> None | HttpResponse:
    response = HttpResponse()

    if not request.user.is_authenticated:
        response.content = "You must be logged in"
        response.status_code = 403
        return response
    return


class AuthoriseView(View):
    def post(self, request: HttpRequest):
        response = HttpResponse()

        returned_value = check_login_info(request.POST)
        if isinstance(returned_value, HttpResponse):
            return returned_value

        username, password = returned_value

        user = authenticate(request, username=username, password=password)
        if user is not None:
            login(request, user)
            response.content = "user logged in"
            response.status_code = 200
            return response
        else:
            response.content = "Authentication issue"
            response.status_code = 401
            return response

    def patch(self, request: HttpRequest):
        response = HttpResponse()
        returned_response = authenticate_request(request)

        if isinstance(returned_response, HttpResponse):
            return returned_response

        if not request.headers.get("password"):
            response.content = "New password not given."
            response.status_code = 400
            return response

        password = request.headers.get("password")
        if self.request.user.check_password(password):
            response.content = "Same password given"
            response.status_code = 409
            return response

        request.user.set_password(password)
        response.content = "Password updated"
        response.status_code = 200
        return response

    def delete(self, request: HttpRequest):
        response = HttpResponse()
        returned_response = authenticate_request(request)

        if isinstance(returned_response, HttpResponse):
            return returned_response

        user = User.objects.get(username=request.user.username)
        logout(request)
        user.delete()

        response.content = "Account removed"
        response.status_code = 200
        return response


class RegisterView(View):
    def post(self, request: HttpRequest):
        response = HttpResponse()
        returned_value = check_login_info(request.POST)
        if isinstance(returned_value, HttpResponse):
            return returned_value

        username, password = returned_value

        user_with_same_username = User.objects.filter(username=username).first()

        if user_with_same_username:
            response.content = "Username already used"
            response.status_code = 409
            return response

        user = User.objects.create(username=username, password=password)
        login(request, user)

        response.content = "New user created"
        response.status_code = 201
        return response


class LogoutView(View):
    def post(self, request: HttpRequest):
        response = HttpResponse()
        returned_response = authenticate_request(request)

        if isinstance(returned_response, HttpResponse):
            return returned_response

        logout(request)
        response.content = "User logged out"
        response.status_code = 200
        return response
