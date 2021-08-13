import React, { useEffect, useState } from "react";
import InputGroup from 'react-bootstrap/InputGroup';
import FormControl from 'react-bootstrap/FormControl';
import { AsyncTypeahead } from 'react-bootstrap-typeahead';
import axios from 'axios';

const NewQuote = ({ quotebookCollection, quotebook, onQuoteeInput, onQuoteTextInput, onQuoteDateInput, onWitnessInput, quoteCount }) => {
    const [isWitnessLoading, setIsWitnessLoading] = useState(false);
    const [witnesses, setWitnesses] = useState([]);
    const [isQuoteeLoading, setIsQuoteeLoading] = useState(false);
    const [quotees, setQuotees] = useState([]);
    const [quoteText, setQuoteText] = useState('');
    const [quoteDate, setQuoteDate] = useState('');

    const backendUrl = window.env.REACT_APP_BACKEND_API
    
    const quoteeRef = React.createRef();
    const witnessRef = React.createRef();

    const handleWitnessSearch = () => {
        setIsWitnessLoading(true);
        axios.get(`${backendUrl}/witnesses/${quotebookCollection}/${quotebook}`)
        .then(res => {
            setWitnesses(res.data);
        })
        setIsWitnessLoading(false);
    };

    const handleQuoteeSearch = () => {
        setIsQuoteeLoading(true);
        axios.get(`${backendUrl}/quotees/${quotebookCollection}/${quotebook}`)
        .then(res => {
            setQuotees(res.data);
        })
        setIsQuoteeLoading(false);
    };

    useEffect(() => {  
        quoteeRef.current.clear()
        witnessRef.current.clear()
        setQuoteText('')
        setQuoteDate('')
    // eslint-disable-next-line
    }, [quoteCount])

    return (
        <tr>
            <td>
                <InputGroup>
                    <AsyncTypeahead
                        id="new-quotee-input-id"
                        ref={quoteeRef}
                        minLength={1}
                        isLoading={isQuoteeLoading}
                        labelKey={(option) => `${option.quoteeName}`}
                        onSearch={handleQuoteeSearch}
                        onChange={(selected) => {
                            if (selected.length > 0) {
                                onQuoteeInput(selected)
                            }
                        }}
                        options={quotees}
                        allowNew
                        newSelectionPrefix="Add new quotee:"
                        placeholder="Quotee"
                    />
                </InputGroup>
            </td>
            <td>
                <FormControl 
                    id="new-quote-text-input-id"
                    name="new-quote-text"
                    placeholder="Quote text..."
                    as="textarea"
                    onChange={(e) => {onQuoteTextInput(e.target.value); setQuoteText(e.target.value)}}
                    value={quoteText}
                />
            </td>
            <td>
                <FormControl 
                    id="new-quote-date-input-id"
                    name="new-quote-date"
                    type="date"
                    onChange={(e) => {onQuoteDateInput(e.target.value); setQuoteDate(e.target.value)}}
                    value={quoteDate}
                />
            </td>
            <td>
                <InputGroup>
                    <AsyncTypeahead
                        id="new-witness-input-id"
                        ref={witnessRef}
                        minLength={1}
                        isLoading={isWitnessLoading}
                        labelKey={(option) => `${option.witnessName}`}
                        onSearch={handleWitnessSearch}
                        onChange={(selected) => {
                            if (selected.length > 0) {
                                onWitnessInput(selected)
                            }
                        }}                        
                        options={witnesses}
                        allowNew
                        newSelectionPrefix="Add new witness:"
                        placeholder="Witness"
                    />
                </InputGroup>
            </td>
        </tr>
    );
}

export default NewQuote;
