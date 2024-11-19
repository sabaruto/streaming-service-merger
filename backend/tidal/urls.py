from django.urls import path

from . import views

urlpatterns = [
    path("", views.index, name="index"),
    path("auth", views.AuthView.as_view(), name="auth")
]