import React from "react";
import Container from 'react-bootstrap/Container';
import { ExclamationDiamond } from 'react-bootstrap-icons';

const ErrorPage = () => {
    return (
        <Container>
            <div>
                <h1>Twas an error, that occurred!</h1>
                <ExclamationDiamond color="red" size={500} />
            </div>
        </Container >
    );
}

export default ErrorPage;
