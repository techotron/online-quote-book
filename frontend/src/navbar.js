import React, { useEffect, useState } from "react";
import { Navbar, Nav, InputGroup, Container, Button } from 'react-bootstrap';
import { Typeahead } from 'react-bootstrap-typeahead';
import { Search } from "react-bootstrap-icons";
import axios from 'axios';

import NewQuoteBookModal from './QuotebookModal/newQuoteBookModal';

const MainNavBar = () => {
    const [quoteBooks, setQuoteBooks] = useState('');
    const [showNewQuoteBookModal, setShowNewQuoteBookModal] = useState(false);

    const backendUrl = window.env.REACT_APP_BACKEND_API

    useEffect(() => {
        axios.get(`${backendUrl}/quotebooks`)
        .then(res => {
            setQuoteBooks(res.data);
        })
        .catch(err => {
            console.error(err);
        })
    }, [])

    return (
        <div>
            <Navbar bg="dark" variant="dark" expand="lg">
                <Container>
                    <Navbar.Brand href="/">Online Quote Book</Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav" />
                    
                    <Nav>
                        <Button variant="outline-light" onClick={() => setShowNewQuoteBookModal(true)}>New Quotebook</Button>
                    </Nav>

                    <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link href="status">Status</Nav.Link>
                    </Nav>
                    </Navbar.Collapse>

                    <Nav className="justify-content-end">
                        <InputGroup>
                            <InputGroup.Prepend>
                                <InputGroup.Text>
                                    <Search size={24} />
                                </InputGroup.Text>
                            </InputGroup.Prepend>
                            <Typeahead
                                id="quotebook-search-id"
                                placeholder="Search quotebooks..."
                                onChange={(selected) => {
                                    if (selected.length > 0) {
                                        window.location.href = `/quotebooks/${selected[0].quoteBookName}`;
                                    }
                                }}
                                options={quoteBooks}
                                labelKey={(option) => `${option.quoteBookName}`}
                            />
                        </InputGroup>
                    </Nav>

                </Container>
            </Navbar>

            <NewQuoteBookModal show={showNewQuoteBookModal} onHide={() => setShowNewQuoteBookModal(false)} />
        </div>
    )
}

export default MainNavBar;
