const breakpoints = [
    '40em', '52em', '64em'
];

const colors = {
    primary: '#0072ff',
    secondary: '#101f34',
    danger: '#ff0043',
    white: '#fff',
    grey: '#9B9B9B',
    paleGrey: '#F4F4F4',
    lightGrey: '#D2D2D2',
    darkGrey: '#4A4A4A',
};

const space = [
    0, 4, 8, 12, 16, 20, 24, 32, 40, 48, 56, 64, 72, 80, 100, 120
];

const fontSizes = [
    12, 14, 18, 24, 28, 38, 56
];

const fontFamilies = {
    sans: '"SkycoinSans"',
    sansBold: '"SkycoinSansBold"',
};

const lineHeights = [
    1, 1.125, 1.25, 1.5
];

const radius = [ 2, 4, 8 ];

export default {
    breakpoints,
    colors,
    space,
    fontSizes,
    fontFamilies,
    lineHeights,
    radius,
    fontLight: '"SkycoinSans"',
    fontBold: '"SkycoinSansBold"',
    container: {
        width: '90%',
        maxWidth: '640px',
    },
    controlHeight: 38,
};
