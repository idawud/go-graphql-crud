import React, { useState, useEffect } from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import axios from 'axios';
import './App.css';

function App() {
	const [ query, setQuery ] = useState('');
	const [ result, setResult ] = useState(null);

	const executeQuery = () => {
		axios
			.post(`http://localhost:8080/graphql`, query)
			.then((res) => {
				setResult(res.data);
			})
			.catch((err) => setResult(err));
	};

	const handleChange = (event) => {
		setQuery(event.target.value);
	};

	return (
		<div className="App">
			<Container>
				<Row>
					<p>HeAER gOES HeRe</p>
				</Row>
				<Row>
					<Col className="input">
						<Form>
							<Form.Control
								className="input"
								as="textarea"
								rows="12"
								placeholder="/* type query here */"
								onChange={handleChange}
							/>
						</Form>
					</Col>
					<Col sm={1}>
						<Button variant="dark" onClick={executeQuery}>
							Execute Query
						</Button>{' '}
					</Col>
					<Col className="result">
						<div>
							<pre>{JSON.stringify(result, null, 2)}</pre>
						</div>
					</Col>
				</Row>
			</Container>
		</div>
	);
}

export default App;
