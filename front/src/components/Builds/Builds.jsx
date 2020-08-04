import React from 'react';
import Card from 'react-bootstrap/Card';
import { useSelector } from 'react-redux';
import {
  SwordEmoji,
  WorkerEmoji,
  ShieldEmoji,
  WrenchEmoji,
} from '../../emojis/'

const Builds = () => {
  const workers = useSelector((state) => state.workers)

  return (
    <Card className="builds">
      <Card.Header>
        <h2>
          <WrenchEmoji />
          Builders
        </h2>
      </Card.Header>
      <Card.Body>
        <Card.Text>
          <SwordEmoji/> <WorkerEmoji/> 
          [{workers.state === "All finished" ? workers.state : workers.sword + " " + workers.state}]
        </Card.Text>
        <Card.Text>
          <ShieldEmoji/> <WorkerEmoji/> 
          [{workers.state === "All finished" ? workers.state : workers.shield + " " + workers.state}]
        </Card.Text>
      </Card.Body>
    </Card>
  );
}

export default Builds;
