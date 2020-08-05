import React, { useState, useEffect } from 'react';
import { Alert, Row, Col, Button } from 'react-bootstrap'
import { useSelector } from 'react-redux';
import {
  FarmerEmoji,
  WorkerEmoji,
} from '../../emojis/'

let finalTime = 0
let done = false
let started = false

const Info = (start) => {
  const info = useSelector((state) => state.info)
  const [time, setTime] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setTime((time) => time + 1)
    }, 1000)

    return () => clearInterval(interval)
  }, [])

  function startWebSocket() {
    start.start()
    setTime(0)
    started = true
  }

 function renderTimer() {
   if (!started) {
     return 0
   }
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
          <p> Failed builds: {info.failedBuilds} </p>
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
      <Button variant="primary" style={{marginLeft: "130px"}} onClick={startWebSocket}>Start</Button>
    </Alert>
  )
}

export default Info;
