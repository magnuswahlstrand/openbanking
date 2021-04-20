import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import Container from 'react-bootstrap/Container';
import Button from 'react-bootstrap/Button';
import Toast from 'react-bootstrap/Toast';
import React, {Component, useState} from 'react';


function copyToClipboard(event, value, alertName) {
    // Uppercase first character
    alertName = alertName.charAt(0).toUpperCase() +
        alertName.slice(1);

    navigator.clipboard.writeText(value).then(
        function () {
            console.log("copied", value);
            //createToast(alertName + " copied.")
        },
        function () {
        }
    );

    event.preventDefault();
    event.stopPropagation();
}


function Section(props) {
    const classes = `mt-5 ${props.faded ? "faded" : ""}`

    return (
        <Container className={classes}>
            <h3>{props.title}</h3>
            {props.children}
        </Container>
    );
};

function PersonalID(props) {
    return (
        <kbd
            className="mr-1"
            style={{cursor: "pointer"}}
            onClick={(e) => copyToClipboard(e, props.id, "personal number")}>
            {props.id}
        </kbd>
    )
}


function DemoToast(props) {
    const [show, setShow] = useState(true);

    return (
        <div
            style={{
                position: 'relative',
                minHeight: '100px',
            }}
        >
            <Toast
                show={show}
                onClose={() => setShow(false)}
                autohide={true}
                delay={10000}
            >
                <Toast.Body>{props.message}</Toast.Body>
            </Toast>
        </div>
    )
}

const list2 = [
    {
        id: 'a',
        name: 'Robin',
    },
    {
        id: 'b',
        name: 'Dennis',
    },
];

class DemoToastList2 extends Component {
    constructor() {
        super();
        this.state = {
            toasts: [],
        }
    }

    render() {
        return (
            <div>
                {this.state.toasts.map((item) => (
                    <DemoToast message={item.name}></DemoToast>
                ))}
            </div>
        )
    }
}


function CardCard(props) {

    return (
        <div className="card">
            <div className="card-header p-0" id="heading_{{index}}">
                <div
                    className="row p-2"
                    data-toggle="collapse"
                    data-target="#collapse_{{index}}"
                    aria-expanded="true"
                    aria-controls="collapse_{{index}}"
                >
                    <div className="col-1">
                        <svg
                            height="48"
                            className="octicon octicon-repo pull-left"
                            viewBox="0 0 12 16"
                            version="1.1"
                            width="36"
                            aria-hidden="true"
                            fill="purple"
                        >
                            <path
                                fill="lightgray"
                                fill-rule="evenodd"
                                d="M4 9H3V8h1v1zm0-3H3v1h1V6zm0-2H3v1h1V4zm0-2H3v1h1V2zm8-1v12c0 .55-.45 1-1 1H6v2l-1.5-1.5L3 16v-2H1c-.55 0-1-.45-1-1V1c0-.55.45-1 1-1h10c.55 0 1 .45 1 1zm-1 10H1v2h2v-1h3v1h5v-2zm0-10H2v9h9V1z"
                            ></path>
                        </svg>
                    </div>

                    <div className="col-5  d-flex align-items-center">
                        <div className="d-flex align-items-start flex-column bd-highlight">
                            <div className="mb-auto"><b>{props.type}</b></div>
                            <div style={{color:"grey"}}>{props.bban}</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}


function App() {
    return (
        <Container className="mt-5 pt-5" onClick={() => list2.concat({name: 'Magnus'})}>
            <Nav className="navbar fixed-top navbar-light bg-light border" id="navbar">
                <Navbar.Brand href="#">
                    <h1 className="mt-2">Open Banking Demo</h1>
                    <p className="lead">by Kyeett</p>
                </Navbar.Brand>
            </Nav>

            <DemoToast message="Balance copied"></DemoToast>

            <Container className="pt-5 content">
                <DemoToastList2 toasts={list2}>

                </DemoToastList2>

                <Section title="0. Choose a demo user">
                    SEB has definied five demo users. Click on one of them to copy, and then proceed to the next
                    step.
                    <br/>

                    <PersonalID id="9311219639"></PersonalID>
                    <PersonalID id="9311219589"></PersonalID>
                    <PersonalID id="8811215477"></PersonalID>
                    <PersonalID id="8811212862"></PersonalID>
                    <PersonalID id="8311211356"></PersonalID>
                </Section>

                <CardCard type="lol" bban="lal"></CardCard>

                <Section title="1. Authorize">
                    Open banking uses Oauth2.0. The link will take you to SEB where you sign in. If you used a real
                    system,
                    it would require BankID (or other form or SCA) to sign in, but in this demo it will directly log you
                    in

                    <center>
                        <Button variant="primary" id="authorize_button" className="mt-3 mb-3">Authorize</Button>
                        {/*onClick="authorize()"*/}
                    </center>
                </Section>


                <Section title="2. Retrieve Token" faded={true}>
                    When the user is authorized, the bank will redirect the browser to an URL specified by the our
                    application and include a query parameter, code=<kbd>&lt;authorization_code&gt;</kbd>.
                    We will now use this code to retrieve a <kbd>Token</kbd> from the bank. Normally, this would be done
                    automatically by the backend.

                    <center>
                        <Button variant="primary" id="authorize_button" className="mt-3 mb-3 disabled">Get
                            Token</Button>
                        {/*onClick="getToken()"*/}
                    </center>
                </Section>


                <Section title="3. Retreive Account Information" faded={true}>
                    Once the <kbd>Token</kbd> is received, we the application can get the account information from the
                    bank
                    server.

                    <center>
                        <Button variant="primary" id="get_accounts_button" className="mt-3 mb-3 disabled">Get
                            Accounts</Button>
                        {/*onClick="getAccounts()"*/}
                    </center>

                    <div className="accordion mt-5" id="accordionAccounts"></div>

                    {/*Quick fix to make UX nicer*/}
                    <div id="get_accounts_padding" className="m-5">&nbsp;</div>
                </Section>

            </Container>
        </Container>
    );
}

export default App;
