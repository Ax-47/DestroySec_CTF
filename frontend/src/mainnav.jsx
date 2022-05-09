import { useEffect } from "react";
import './material/css/mainnav.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import Navbar from 'react-bootstrap/Navbar';
//import NavbarBrand from 'react-bootstrap/esm/NavbarBrand';
import Container from 'react-bootstrap/Container';
//import NavbarToggle from 'react-bootstrap/esm/NavbarToggle';
//import { Offcanvas } from 'bootstrap';
//import OffcanvasHeader from 'react-bootstrap/esm/OffcanvasHeader';
import Nav from 'react-bootstrap/Nav';
//import NavLink from 'react-bootstrap/esm/NavLink';
//import NavDropdown from 'react-bootstrap/NavDropdown';
//import Form from 'react-bootstrap/Form'
//import FormControl from 'react-bootstrap/FormControl'
import Button from 'react-bootstrap/esm/Button';
import {LinkContainer} from 'react-router-bootstrap'
import FormSelect from 'react-bootstrap/esm/FormSelect';
//import Select from 'react-select'
//import { useInRouterContext } from 'react-router-dom';
import { Link } from "react-router-dom";

import { useTranslation } from "react-i18next";
import i18next from "i18next";






function  Navapp() {

  
  /*
  const lang = [
    { value: "en", label: "en"},
    { value: "th", label: "th"}
  ]*/

  
    const { i18n, t } = useTranslation(["mainnav"]);  

    useEffect(() => {
      if (localStorage.getItem("i18nextLng")?.length > 2) {
        i18next.changeLanguage("en");
      }
    }, []);
  
    const handleLanguageChange = (e) => {
      i18n.changeLanguage(e.target.value);
    };

    return (
        <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark" className='navh'>
        <Container>
        <LinkContainer to="/">
        <Navbar.Brand >DestroySec   CTF</Navbar.Brand>
        </LinkContainer>
        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse id="responsive-navbar-nav">
          <Nav className="me-auto">
            <LinkContainer to="/">
            <Nav.Link>{t("home")}</Nav.Link>
            </LinkContainer>
            <LinkContainer to="/features">
            <Nav.Link>{t("features")}</Nav.Link>
            </LinkContainer>
            <LinkContainer to="/pricing">
            <Nav.Link >{t("pricing")}</Nav.Link>
            </LinkContainer>
            
          </Nav>
          <Nav className='me-3'>
            <Link to="/registerandlogin">
            <Button variant="outline-warning" >{t("sign up sign in")}</Button>
            </Link>
          </Nav>
          <br/>
          <Nav>
            <FormSelect className="bg-dark border-0 text-white" value={localStorage.getItem("i18nextLng")} onChange={handleLanguageChange}  >
              <option value="en">en</option>
              <option value="th">th</option>
            </FormSelect>
           
          </Nav>
        </Navbar.Collapse>
        </Container>
      </Navbar>
 

   
    )
}

export default Navapp;