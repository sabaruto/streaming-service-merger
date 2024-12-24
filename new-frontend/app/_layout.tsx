import {type ErrorBoundaryProps, Stack} from "expo-router";
import {ThemeProvider} from "@react-navigation/native";
import {useColorScheme} from "@/app-example/hooks/useColorScheme";
import {Themes} from "@/constants/Themes";
import {Text, View} from "react-native";

export function ErrorBoundary({error, retry}: ErrorBoundaryProps) {
    return (
        <View style={{flex: 1, backgroundColor: "red"}}>
            <Text>{error.message}</Text>
            <Text onPress={retry}>Try Again?</Text>
        </View>
    );
}

export default function RootLayout() {
    let theme = useColorScheme() === "dark" ? Themes["dark"] : Themes["light"]

    return (
        <ThemeProvider value={theme}>
            <Stack screenOptions={{headerShown: false}}/>
        </ThemeProvider>
    );
}
