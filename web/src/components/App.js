import React from 'react';
import { Switch, Route, BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';
import styled, { injectGlobal, ThemeProvider } from 'styled-components';
import { Flex } from 'grid-styled';

import Routes from './Routes';

import theme from './theme';
import store from '../store';

import skycoinLight from '../fonts/skycoin-regular-webfont.woff';
import skycoinBold from '../fonts/skycoin-bold-webfont.woff';

injectGlobal`
  @font-face {
    font-family: ${theme.fontLight};
    src: url(${skycoinLight}) format('woff');
    font-weight: normal;
  }
  
    @font-face {
    font-family: ${theme.fontBold};
    src: url(${skycoinBold}) format('woff');
    font-weight: normal;
  }
  
  * {
    box-sizing: border-box;
  }
    
  body {
    margin: 0;
    padding: 0;
    font-family: ${theme.fontLight}, Arial, Helvetica, sans-serif;
    color: ${theme.colors.darkGrey};
  }
  
  ul {
    margin: 0;
    padding: 0;
  }
  a {
    color: ${theme.colors.primary};
    text-decoration: none;
  }
`;

const Wrapper = styled(Flex) `
    min-height: 100vh;
    max-width: 100%;
    overflow-x: hidden;
`;

const Root = ({ locale, ...props }) => (
    <ThemeProvider theme={theme}>
        <Wrapper flexDirection="column">
            <Routes {...props} />
        </Wrapper>
    </ThemeProvider>
);

export default () => (
    <Provider store={store}>
        <BrowserRouter>
            <Switch>
                <Route path="/" render={props => <Root {...props} locale="en" />} />
            </Switch>
        </BrowserRouter>
    </Provider>
);
