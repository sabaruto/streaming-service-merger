from django.urls import path
from authorisation import views

urlpatterns = [
    path("", views.AuthoriseView.as_view(), name="user"),
    path("csrf/", views.CSRFView.as_view(), name="csrf"),
    path("register/", views.RegisterView.as_view(), name="register"),
    path("logout/", views.LogoutView.as_view(), name="logout"),
]
