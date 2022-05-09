import React from 'react';
//import './App.css';
import Navapp from './mainnav';
import Home from './material/page/index/index'
import Features from './material/page/features';
import Pricing from './material/page/pricing';
import RegisLog from './material/page/regisandlog/login';
import RegisSign from './material/page/regisandlog/signup';
import { Routes, Route, BrowserRouter } from 'react-router-dom';



function App() {
    return ( 
    <BrowserRouter >
        <Navapp />

        <Routes>
        <Route exact path = '/'element = { <Home /> } > 
        </Route> <Route path = '/features'element = { <Features/> } > 
        </Route> 
        <Route path = '/pricing'element = { <Pricing/> } >
        </Route> 
        <Route path = '/login'element = { <RegisLog/> } > </Route> 
        <Route path = '/register'element = { <RegisSign/> } > </Route> 
        </Routes >

      </BrowserRouter>
    );
}

export default App;