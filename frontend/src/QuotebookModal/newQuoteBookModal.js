import React, { useState } from "react";
import { Modal, Button, Form } from 'react-bootstrap';
import axios from 'axios';

const NewQuoteBookModal = ({ show, onHide }) => {
    const [newQuotebookName, setNewQuotebookName] = useState('');
    const [newQuotebookCollection, setNewQuotebookCollection] = useState('public');

    const backendUrl = window.env.REACT_APP_BACKEND_API

    const onSubmit = (quotebookCollection, quotebookName) => {
        axios.post(`${backendUrl}/quotebook/${quotebookCollection}/${quotebookName}`)
        .then(res => {
            window.location.href = `/quotebooks/${quotebookCollection}/${quotebookName}`
        })
        .catch(err => {
            console.error(err)
        })
    }

    const handleQuotebookNameChange = (e) => {
        setNewQuotebookName(e.target.value)
    }

    const handleQuotebookCollectionChange = (e) => {
        setNewQuotebookCollection(e.target.value)
    }

    return (
        <Modal animation={true} show={show} onHide={onHide}>
            <Modal.Header closeButton>
                <Modal.Title>New Quotebook</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form>
                    <b>Quotebook Collection</b>
                    <Form.Group className="pt-2">
                        <Form.Control type="text" placeholder="public" id="new-quotebook-collection-text-box" disabled={true} onChange={(e) => handleQuotebookCollectionChange(e)} />
                    </Form.Group>
                </Form>
                <Form>
                    <b>Quotebook Name</b>
                    <Form.Group className="pt-2">
                        <Form.Control type="text" placeholder="Your new quotebook" id="new-quotebook-text-box" onChange={(e) => handleQuotebookNameChange(e)} />
                    </Form.Group>
                </Form>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="primary" type="submit" id="new-quotebook-button" onClick={(e) => onSubmit(newQuotebookCollection, newQuotebookName)}>Create</Button>
            </Modal.Footer>
        </Modal>
    )
}

export default NewQuoteBookModal;
