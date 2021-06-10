import React, { useEffect, useState } from "react";
import Container from 'react-bootstrap/Container';
import Spinner from 'react-bootstrap/Spinner';
import Table from 'react-bootstrap/Table';
import axios from 'axios';

const Info = () => {
  const [backendInfo, setBackendInfo] = useState('');
  const [isLoaded, setIsLoaded] = useState(false);
  const backendUrl = window.env.REACT_APP_BACKEND_API

  useEffect(() => {
    axios.get(`${backendUrl}/info`)
      .then(res => {
        setBackendInfo(res.data);
        setIsLoaded(true)
      })
  }, [backendUrl])

  return (
    <Container>
      {isLoaded ?
        <Container>
          <h1>Info page</h1>
          <b>DB Schema: </b>{backendInfo.db_version}
          <br />
          <b>Is Dirty: </b>{backendInfo.db_dirty.toString()}
          <br />
          <br />
          <br />
          <Table striped bordered hover>
            <thead>
              <tr>
                <th>Setting</th>
                <th>Value</th>
              </tr>
            </thead>
            <tbody>
              {
                Object.entries(backendInfo).map((key) =>
                  <tr>
                    <td>{key[0].toString()}</td>
                    <td>{key[1].toString()}</td>
                  </tr>
                )
              }
            </tbody>
          </Table>
        </Container>
        : <Spinner animation="border" variant="primary" />}
    </Container>
  );
}

export default Info;
