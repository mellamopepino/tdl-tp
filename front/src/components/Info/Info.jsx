import React, { useState, useEffect } from 'react';
import { Alert, Row, Col } from 'react-bootstrap'
import { useSelector } from 'react-redux';
import {
  FarmerEmoji,
  WorkerEmoji,
} from '../../emojis/'

let finalTime = 0
let done = false

const Info = () => {
  const info = useSelector((state) => state.info)
  const [time, setTime] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setTime((time) => time + 1)
    }, 1000)

    return () => clearInterval(interval)
  }, [])

 function renderTimer() {
   if (done) {
     return finalTime
   }
   if (info.done) {
     done = true
     finalTime = time
     return finalTime
   }
   return time
 }
  
  return (
    <Alert variant="info">
      <Row>
        <Col className="d-flex justify-content-center">
          <h1 className="align-self-center">
            Time {renderTimer()}
          </h1>
        </Col>
        <Col className="mx-2">
          <h4>Jobs</h4>
          <p> In progress: {info.jobsInProgress} </p>
          <p> Finished: {info.jobsFinished} </p>
        </Col>
        <Col className="mx-2">
          <h4>Total workers</h4>
          <p>
            <FarmerEmoji/> [{info.totalGatherers}]
          </p>
          <p>
            <WorkerEmoji/> [{info.totalWorkers}]
          </p>
        </Col>
      </Row>
    </Alert>
  )
}

export default Info;
