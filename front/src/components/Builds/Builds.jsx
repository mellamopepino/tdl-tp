import React from 'react';
import Card from 'react-bootstrap/Card';
import {
  SwordEmoji,
  WorkerEmoji,
  ShieldEmoji,
  WrenchEmoji,
} from '../../emojis/'

const Builds = () => {

  return (
    <Card className="builds">
      <Card.Header>
        <h2>
          <WrenchEmoji />
          Builds
        </h2>
      </Card.Header>
      <Card.Body>
        <Card.Text>
          <p>
            <SwordEmoji/> <WorkerEmoji/> [2 working...]
          </p>
          <p>
            <ShieldEmoji/> <WorkerEmoji/> [1 working...]
          </p>
        </Card.Text>
      </Card.Body>
    </Card>
  );
}

export default Builds;
