import {View, ViewProps, ViewStyle} from "react-native";
import {themedFuncToStyle, ThemeProps} from "@/constants/Themes";

export type ThemedViewProps = ViewProps & ThemeProps<ViewStyle>

export function ThemedView({theme, style, themedFns, ...otherProps}: ThemedViewProps) {
    let themedStyle = themedFns ? themedFuncToStyle(theme, themedFns) : []
    return <View style={[
        style,
        {backgroundColor: theme.colors.background},
        themedStyle
    ]} {...otherProps}/>
}
