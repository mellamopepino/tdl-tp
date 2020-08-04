import React from 'react';
import Card from 'react-bootstrap/Card';
import { useSelector } from 'react-redux';
import {
  TreeEmoji,
  FarmerEmoji,
  StoneEmoji,
  AxeEmoji,
  MedalEmoji,
} from '../../emojis/'

const Resources = () => {
  const gatherers = useSelector((state) => state.gatherers)
  const resources = useSelector((state) => state.resources)

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
          [{resources.wood}] <TreeEmoji/> <FarmerEmoji/> 
          [{gatherers.state === "All finished" ? gatherers.state : gatherers.wood + " " + gatherers.state}]
        </Card.Text>
        <Card.Text>
        [{resources.stone}] <StoneEmoji/> <FarmerEmoji/> 
        [{gatherers.state === "All finished" ? gatherers.state : gatherers.stone + " " + gatherers.state}]
        </Card.Text>
        <Card.Text>
        [{resources.gold}] <MedalEmoji/> <FarmerEmoji/> 
        [{gatherers.state === "All finished" ? gatherers.state : gatherers.gold + " " + gatherers.state}]
        </Card.Text>
      </Card.Body>
    </Card>
  );
}

export default Resources;
