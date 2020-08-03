import React from 'react';
import Card from 'react-bootstrap/Card';
import { useSelector } from 'react-redux';
import {
  TreeEmoji,
  FarmerEmoji,
  StoneEmoji,
  AxeEmoji,
} from '../../emojis/'

const Resources = () => {
  const gatherers = useSelector((state) => state.gatherers)

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
          <TreeEmoji/> <FarmerEmoji/> [{gatherers.wood} working...]
        </Card.Text>
        <Card.Text>
          <StoneEmoji/> <FarmerEmoji/> [{gatherers.stone} working...]
        </Card.Text>
        <Card.Text>
          <StoneEmoji/> <FarmerEmoji/> [{gatherers.gold} working...]
        </Card.Text>
      </Card.Body>
    </Card>
  );
}

export default Resources;
