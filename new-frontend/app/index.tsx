import {Suspense} from "react";
import {Text, View} from "react-native";
import {center, fullScreen, primaryBackground} from "@/constants/Styles";
import {useTheme} from "@react-navigation/native";
import {LocalAuthComponent} from "@/components/forms/localAuth";
import {ThemedView} from "@/components/themedView";

// function LazyComponent() {
//     const _ = usePromise(() => new Promise((resolve) => setTimeout(resolve, 1000)), []);
//     return <Text>Lazy is loaded</Text>;
// }

export default function Index() {
    let theme = useTheme()

    return (
        // Add Router here
        <ThemedView
            theme={theme}
            themedFns={primaryBackground}
            style={[center, fullScreen]}
        >
            <Suspense fallback={<View><Text>Loading</Text></View>}>
                <LocalAuthComponent/>
            </Suspense>
        </ThemedView>
    );
}

