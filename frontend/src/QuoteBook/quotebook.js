import React, { useEffect, useState } from "react";
import Container from 'react-bootstrap/Container';
import Spinner from 'react-bootstrap/Spinner';
import Table from 'react-bootstrap/Table';
import axios from 'axios';
import dayjs from 'dayjs';

import ErrorPage from '../Errors/errorpage';

const QuoteBook = ({ location }) => {
    const [quotes, setQuotes] = useState('');
    const [isLoaded, setIsLoaded] = useState(false);
    const [quotebookCollection, setQuotebookCollection] = useState('');
    const [quotebook, setQuoteBook] = useState('');
    const [isError, setIsError] = useState(false);

    const backendUrl = window.env.REACT_APP_BACKEND_API

    useEffect(() => {
        const pathname = location.pathname.split('/');
        setQuoteBook(pathname[pathname.length - 1]);
        setQuotebookCollection(pathname[pathname.length - 2]);
        
        if (quotebook !== '') {
            axios.get(`${backendUrl}/quotes/${quotebookCollection}/${quotebook}`)
            .then(res => {
                setQuotes(res.data);
                setIsLoaded(true);
                setIsError(false);
            })
            .catch(err => {
                console.error(err);
                setIsError(true);
            })
        }
    }, [backendUrl, location.pathname, quotebook, quotebookCollection])

    const dateFormatter = (date) => {
        let fdate = dayjs(date).format('DD-MM-YYYY')
        return fdate
    }

    return (
        <Container>
            {(isLoaded && !isError) ?
            <Container>
                <h1>Quotes: {quotebook}</h1>
                <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>Quotee</th>
                        <th>Quote</th>
                        <th>Quote Date</th>
                        <th>Witness</th>
                    </tr>
                </thead>
                <tbody>
                    {
                    Object.entries(quotes).map((key) =>
                        <tr>
                            <td>{key[1].quotee}</td>
                            <td>{key[1].quoteText}</td>
                            <td>{dateFormatter(key[1].quoteDate)}</td>
                            <td>{key[1].witness}</td>
                        </tr>
                    )
                    }
                </tbody>
                </Table>
            </Container>
            : isError ? <ErrorPage /> : <Spinner animation="border" variant="primary" />}
        </Container>
    );
}

export default QuoteBook;
