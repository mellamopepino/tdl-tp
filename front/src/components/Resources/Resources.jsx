import React from 'react';
import Card from 'react-bootstrap/Card';
import {
  TreeEmoji,
  FarmerEmoji,
  StoneEmoji,
  AxeEmoji,
} from '../../emojis/'

const Resources = () => {

  return (
    <Card className="resources">
      <Card.Header>
        <h2>
          <AxeEmoji />
          Resources
        </h2>
      </Card.Header>
      <Card.Body>
        <Card.Text>
          <p>
            <TreeEmoji/> [3] <FarmerEmoji/> [2 working...]
          </p>
          <p>
            <StoneEmoji/> [2] <FarmerEmoji/> [1 working...]
          </p>
        </Card.Text>
      </Card.Body>
    </Card>
  );
}

export default Resources;
