import {StyleProp, Text, TextInput, TextInputProps, TextProps, TextStyle} from "react-native";
import {FontSize, themedFontSizeStyle, themedFuncToStyle, ThemeProps} from "@/constants/Themes";

interface ThemedTextableProps extends ThemeProps<TextStyle> {
    fontSize?: FontSize,
    style?: StyleProp<TextStyle>
}

export type ThemedTextType = TextProps & ThemedTextableProps
export type ThemedTextInputType = TextInputProps & ThemedTextableProps

export function createThemeTextStyles(props: ThemedTextableProps) {

    let fontStyle = themedFontSizeStyle(props.theme, props.fontSize ? props.fontSize : "regular")
    let themedStyle = props.themedFns ? themedFuncToStyle(props.theme, props.themedFns) : []
    let defaultStyle = {
        fontFamily: fontStyle.fontFamily,
        fontWeight: fontStyle.fontWeight,
        color: props.theme.colors.text
    }

    return {defaultStyle, themedStyle}
}

export function ThemedText({theme, style, fontSize, themedFns, ...otherProps}: ThemedTextType) {
    let {
        defaultStyle,
        themedStyle
    } = createThemeTextStyles({theme, style, fontSize, themedFns: themedFns})


    return <Text style={[
        style,
        defaultStyle,
        themedStyle,
    ]} {...otherProps}/>
}

export function ThemedTextInput({theme, style, fontSize, themedFns, ...otherProps}: ThemedTextInputType) {
    let {
        defaultStyle,
        themedStyle,
    } = createThemeTextStyles({theme, style, fontSize, themedFns: themedFns})

    return <TextInput style={[
        defaultStyle,
        themedStyle,
        style
    ]} {...otherProps} />
}
