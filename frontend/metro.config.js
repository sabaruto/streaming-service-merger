const { getDefaultConfig } = require("expo/metro-config");

module.exports = (() => {
    const config = getDefaultConfig(__dirname);

    const {
        transformer,
        resolver: { assetExts, sourceExts },
    } = config;

    config.transformer = {
        ...transformer,
        babelTransformerPath: require.resolve("react-native-sass-transformer"),
    };
    config.resolver = {
        assetExts: assetExts,
        sourceExts: [...sourceExts, "scss", "sass"],
    };

    return config;
})();
