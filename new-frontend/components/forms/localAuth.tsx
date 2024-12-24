import {Button} from "react-native"
import Form from "@/components/form";
import {useTheme} from "@react-navigation/native";
import {footer, mainCard} from "@/constants/Styles";
import {MutableRefObject, useState} from "react";
import {GestureResponderEvent} from "react-native/Libraries/Types/CoreEventTypes";
import {ThemedView} from "@/components/themedView";
import {ThemedText} from "@/components/themedText";

type formText = {
    title: string
    footer: string,
    footerButton: string,
}

type LocalAuthProps = {
    text: formText
    username: {
        callback: (text: string) => void
        value: string
    }
    password: {
        callback: (text: string) => void
        value: string
    }
    button: {
        submit: (event: GestureResponderEvent) => void
        reset: (event: GestureResponderEvent) => void
    }
}

const registerText: formText = {
    title: "Register",
    footer: "Already have user?",
    footerButton: "Log in"
}

const signInText: formText = {
    title: "Log In",
    footer: "New User?",
    footerButton: "Register"
}


export function LocalAuthComponent() {
    const [isRegister, setShowRegister] = useState(true)
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const currentText = isRegister ? registerText : signInText

    function onTextChange(ref: MutableRefObject<string>) {
        return (text: string) => {
        }
    }

    function onSubmit(event: GestureResponderEvent) {
        return
    }

    return LocalAuthForm({
        text: currentText,
        username: {
            callback: setUsername,
            value: username
        },
        password: {
            callback: setPassword,
            value: password
        },
        button: {
            submit: onSubmit,
            reset: () => {
                setShowRegister(!isRegister)
                setUsername("")
                setPassword("")
            }
        }
    })
}

function LocalAuthForm(props: LocalAuthProps) {
    let theme = useTheme()

    const form = Form({
        header: props.text.title,
        inputs: [
            {
                text: "Username",
                value: props.username.value,
                callback: props.username.callback
            },
            {
                text: "Password",
                value: props.password.value,
                callback: props.password.callback
            }
        ],
        submitCallback: props.button.submit
    })
    let Value = ThemedText
    return (
        <ThemedView theme={theme} themedFns={mainCard}>
            {form}
            <ThemedView theme={theme} style={footer}>
                <ThemedText theme={theme}>{props.text.footer}</ThemedText>
                <Button title={props.text.footerButton} onPress={props.button.reset}/>
            </ThemedView>
        </ThemedView>
    )
}