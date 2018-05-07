const breakpoints = [
    '40em', '52em', '64em'
];

const colors = {
    base: '#0072ff',
    dark: '#101f34',
    white: '#fff',
    lightGrey: '#b3bdca',
    black: '#000',
};

const space = [
    0, 4, 8, 12, 16, 20, 24, 32, 40, 48, 56, 64, 72, 80, 100, 120
];

const fontSizes = [
    12, 14, 18, 24, 28, 38, 56
];

const fontFamilies = {
    sans: '"skycoin-light", sans-serif',
    sansBold: '"skycoin-bold", sans-serif',
};

const lineHeights = [
    1, 1.125, 1.25, 1.5
];

export default {
    breakpoints,
    colors,
    space,
    fontSizes,
    fontFamilies,
    lineHeights,
    fontLight: '\'SkycoinSans\'',
    fontBold: '\'SkycoinSansBold\'',
    container: {
        width: '90%',
        maxWidth: '1280px',
    }
};
