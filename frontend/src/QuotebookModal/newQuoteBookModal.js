import { Modal, Button } from 'react-bootstrap';

const NewQuoteBookModal = ({ show, onHide }) => {
    return (
        <Modal animation={true} show={show} onHide={onHide}>
            <Modal.Header closeButton>
                <Modal.Title>New Quotebook</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>Modal body text goes here.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="secondary">Close</Button>
                <Button variant="primary">Save changes</Button>
            </Modal.Footer>
        </Modal>
    )
}

export default NewQuoteBookModal;
