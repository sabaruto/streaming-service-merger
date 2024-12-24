import {useTheme} from "@react-navigation/native";
import React from "react";
import {ThemedView} from "@/components/themedView";
import {ThemedText, ThemedTextInput} from "@/components/themedText";
import {GestureResponderEvent} from "react-native/Libraries/Types/CoreEventTypes";
import {center, defaultHeader, form, h1, input} from "@/constants/Styles";
import {buttonText, ThemedPressable} from "@/components/ThemedPressable";

type FormProps = {
    header: string
    inputs: FormInputProps[]
    submitCallback: (event: GestureResponderEvent) => void
}

export default function Form(props: FormProps) {
    let theme = useTheme()
    return (
        <ThemedView theme={theme} style={[center, form, {alignItems: "stretch"}]}>
            <ThemedText
                style={h1}
                themedFns={defaultHeader}
                fontSize={"bold"}
                theme={theme}>
                {
                    props.header
                }
            </ThemedText>
            {props.inputs.map((props, index) => FormInput(props, index))}
            <ThemedPressable theme={theme} style={center}>
                <ThemedText theme={theme} themedFns={buttonText}>Submit</ThemedText>
            </ThemedPressable>
        </ThemedView>
    )
}

type FormInputProps = {
    text: string
    value: string
    callback: (text: string) => void
}

export function FormInput(props: FormInputProps, id?: number) {
    let theme = useTheme()
    return (
        <ThemedView theme={theme} key={id}>
            <ThemedText theme={theme}>
                {props.text}
            </ThemedText>
            <ThemedTextInput themedFns={input} theme={theme} value={props.value} onChangeText={props.callback}/>
        </ThemedView>

    )
}
