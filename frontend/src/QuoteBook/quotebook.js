import React, { useEffect, useState } from "react";
import Container from 'react-bootstrap/Container';
import Spinner from 'react-bootstrap/Spinner';
import Button from 'react-bootstrap/Button';
import Table from 'react-bootstrap/Table';
import Form from 'react-bootstrap/Form';
import axios from 'axios';
import dayjs from 'dayjs';

import ErrorPage from '../Errors/errorpage';
import NewQuote from "../NewQuote/newquote";

const QuoteBook = ({ location }) => {
    const [quotes, setQuotes] = useState('');
    const [isLoaded, setIsLoaded] = useState(false);
    const [quotebookCollection, setQuotebookCollection] = useState('');
    const [quotebook, setQuoteBook] = useState('');
    const [isError, setIsError] = useState(false);
    const [newQuoteQuotee, setNewQuoteQuotee] = useState('');
    const [newQuoteText, setNewQuoteText] = useState('');
    const [newQuoteDate, setNewQuoteDate] = useState('');
    const [newQuoteWitness, setNewQuoteWitness] = useState('');
    const [quoteCount, setQuoteCount] = useState(0);

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
                setQuoteCount(res.data.length)
            })
            .catch(err => {
                console.error(err);
                setIsError(true);
            })
        }
    }, [backendUrl, location.pathname, quotebook, quotebookCollection, quoteCount])

    const handleNewQuotee = (quotee) => {
        if (quotee.length > 0) {
            var quoteeName = quotee[0].customOption ? quotee[0].label : quotee[0].quoteeName
            setNewQuoteQuotee(quoteeName)
        }
    }

    const handleNewWitness = (witness) => {
        if (witness.length > 0) {
            var witnessName = witness[0].customOption ? witness[0].label : witness[0].witnessName
            setNewQuoteWitness(witnessName)
        }
    }    

    const handleSubmit = (e) => {
        e.preventDefault()
        var payload = {
            quotee: newQuoteQuotee,
            quoteText: newQuoteText,
            quoteDate: newQuoteDate,
            witness: newQuoteWitness
        }
        axios.post(`${backendUrl}/quotes/${quotebookCollection}/${quotebook}`, payload)
        .catch(err => {
            console.error(err);
            if (err.response.status === 400) {
                // TODO: Create alarm popup to feed this message back to users
                console.log(`Click on "Add new witness/quotee to add a row to the DB"`)
            }
        })

        setQuoteCount(quoteCount + 1)
    }

    const dateFormatter = (date) => {
        let fdate = dayjs(date).format('DD-MM-YYYY')
        return fdate
    }

    return (
        <Container fluid={true}>
            {(isLoaded && !isError) ?
            <Container fluid={true}>
                <h1>Quotes: {quotebook}</h1>
                Total Quotes: {quoteCount}
                <Table striped bordered hover responsive="xl">
                    <thead>
                        <tr>
                            {["Quotee", "Quote", "Quote Date", "Witness"].map((header) => (
                                <th>{header}</th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {Object.entries(quotes).map((key) =>
                            <tr>
                                <td>{key[1].quotee}</td>
                                <td style={{whiteSpace:'pre-wrap', wordWrap: 'break-word'}}>{key[1].quoteText}</td>
                                <td>{dateFormatter(key[1].quoteDate)}</td>
                                <td>{key[1].witness}</td>
                            </tr>)}
                            <NewQuote 
                                quotebookCollection={quotebookCollection}
                                quotebook={quotebook}
                                onQuoteeInput={handleNewQuotee}
                                onQuoteTextInput={setNewQuoteText}
                                onQuoteDateInput={setNewQuoteDate}
                                onWitnessInput={handleNewWitness}
                                quoteCount={quoteCount}
                            />
                    </tbody>
                </Table>
                <Form onSubmit={(e) => handleSubmit(e)}>
                    <Button type="submit">Add Quote</Button>
                </Form>
            </Container>
            : isError ? <ErrorPage /> : <Spinner animation="border" variant="primary" />}
        </Container>
    );
}

export default QuoteBook;
