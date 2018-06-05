import React from 'react';
import { Switch, Route } from 'react-router-dom';
import { Provider } from 'react-redux';
import { ConnectedRouter } from 'react-router-redux';
import styled, { injectGlobal, ThemeProvider } from 'styled-components';
import { Flex } from 'grid-styled';

import Routes from './Routes';

import theme from './theme';
import store, { history } from '../store';

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
  
  html {
    -webkit-font-smoothing: antialiased;
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
`;

const Root = ({ locale, ...props }) => (
    <ThemeProvider theme={theme}>
        <Wrapper flexDirection="column" justifyContent="center">
            <Routes {...props} />
        </Wrapper>
    </ThemeProvider>
);

export default () => (
    <Provider store={store}>
        <ConnectedRouter history={history}>
            <Switch>
                <Route path="/" render={props => <Root {...props} locale="en" />} />
            </Switch>
        </ConnectedRouter>
    </Provider>
);
