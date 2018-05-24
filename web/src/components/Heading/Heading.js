import styled from 'styled-components';
import { space, width, fontSize, lineHeight, color, textAlign } from 'styled-system';
import createComponentFromTagProp from 'react-create-component-from-tag-prop';

const Heading = createComponentFromTagProp({
    tag: 'h2',
    prop: 'as',
    propsToOmit: ['fontSize', 'color', 'bg', 'mt', 'mb', 'my', 'heavy', 'caps', 'textAlign'],
});

export default styled(Heading)`   
    margin: 0;         
    font-weight: normal;
    
    ${width}
    ${space}
    ${fontSize}
    ${lineHeight}
    ${color}
    ${textAlign}
`;
