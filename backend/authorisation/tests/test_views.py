from django.contrib.auth.models import User
from django.test import TestCase


class UserTestCase(TestCase):

    def setUp(self):
        self.authorised_user = User.objects.create_user(
            username="test", password="test"
        )
        self.slug: str

    def no_info_test(self):
        response = self.client.post(self.slug)
        self.assertEqual(response.status_code, 400)

    def missing_info_test(self):
        response = self.client.post(self.slug, {"username": "test"})
        self.assertEqual(response.status_code, 400)

        response = self.client.post(self.slug, {"password": "test"})
        self.assertEqual(response.status_code, 400)


class AuthoriseTests(UserTestCase):
    def setUp(self):
        super().setUp()
        self.slug = "/auth/"

    def test_no_info(self):
        self.no_info_test()

    def test_missing_info(self):
        self.missing_info_test()

    def test_incorrect_info(self):
        response = self.client.post(self.slug, {"username": "bob", "password": "bob"})
        self.assertEqual(response.status_code, 401)

    def test_correct_info(self):
        response = self.client.post(
            self.slug, {"username": "test", "password": "test"}, secure="true"
        )
        self.assertEqual(response.status_code, 200)


class LogoutTests(UserTestCase):
    def setUp(self):
        super().setUp()
        self.slug = "/auth/logout/"

    def test_logged_out_user(self):
        response = self.client.post(self.slug, secure="true")
        self.assertEqual(response.status_code, 403)

    def test_successful_logout(self):
        self.client.login(username="test", password="test")

        response = self.client.post(self.slug, secure="true")
        self.assertEqual(response.status_code, 200)

        response = self.client.post(self.slug, secure="true")
        self.assertEqual(response.status_code, 403)

        login_result = self.client.login(username="test", password="test")
        self.assertTrue(login_result)


class CreateUserTests(UserTestCase):
    def setUp(self):
        super().setUp()
        self.slug = "/auth/register/"

    def test_no_info(self):
        self.no_info_test()

    def test_missing_info(self):
        self.missing_info_test()

    def test_existing_user(self):
        response = self.client.post(
            self.slug, {"username": "test", "password": "test"}, secure="true"
        )
        self.assertEqual(response.status_code, 409)

    def test_new_user(self):
        response = self.client.post(
            self.slug, {"username": "user", "password": "user"}, secure="true"
        )
        self.assertEqual(response.status_code, 201)

        # Check new user is created
        user = User.objects.filter(username="user").first()
        self.assertIsNotNone(user)


class DeleteUserTests(UserTestCase):
    def setUp(self):
        super().setUp()
        self.slug = "/auth/"

    def test_logged_out_user(self):
        response = self.client.delete(self.slug, secure="true")
        self.assertEqual(response.status_code, 403)

    def test_remove_user(self):
        self.client.login(username="test", password="test")

        response = self.client.delete(self.slug, secure="true")
        self.assertEqual(response.status_code, 200)

        login_result = self.client.login(username="test", password="test")
        self.assertFalse(login_result)

        response = self.client.delete(self.slug, secure="true")
        self.assertEqual(response.status_code, 403)


class UserPasswordTests(UserTestCase):
    def setUp(self):
        super().setUp()
        self.slug = "/auth/"

    def test_no_info(self):
        self.client.login(username="test", password="test")
        response = self.client.patch(self.slug, secure="true")
        self.assertEqual(response.status_code, 400)

    def test_logged_out_user(self):
        response = self.client.patch(
            self.slug, headers={"password": "password"}, secure="true"
        )
        self.assertEqual(response.status_code, 403)

    def test_same_password(self):
        self.client.login(username="test", password="test")

        response = self.client.patch(
            self.slug, headers={"password": "test"}, secure="true"
        )
        self.assertEqual(response.status_code, 409)

    def test_new_password(self):
        self.client.login(username="test", password="test")

        response = self.client.patch(
            self.slug, headers={"password": "password"}, secure="true"
        )
        self.assertEqual(response.status_code, 200)
