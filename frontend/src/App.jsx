import React, { Suspense } from 'react';
//import './App.css';
import Navapp from './mainnav';
import Home from './material/page/index/index'
import Features from './material/page/features';
import Pricing from './material/page/pricing';

import RegisSign from './material/page/regisandlog/signup';
import { Routes, Route, BrowserRouter } from 'react-router-dom';

const LazyLogin = React.lazy(() => import('./material/page/regisandlog/login'))
const LazyRegis = React.lazy(() => import('./material/page/regisandlog/signup'))
const LazyHome = React.lazy(() => import('./material/page/index/index'))
const LazyFeatures = React.lazy(() => import('./material/page/features'))
const LazyPricing = React.lazy(() => import('./material/page/pricing'))

function App() {
    return ( 
    <BrowserRouter >
    <Suspense>
        <Navapp />
        </Suspense>

        <Routes>
        <Route exact path='/' element = { <React.Suspense><Home /></React.Suspense> } > 
        </Route> <Route path='/features' element = { <React.Suspense><Features/> </React.Suspense>} > 
        </Route> 
        <Route path='/pricing' element = {<React.Suspense><Pricing/></React.Suspense> } >
        </Route> 
        <Route  path='/login' element = {<React.Suspense fallback={<div>loading</div>}><LazyLogin/> </React.Suspense>} > </Route> 
        <Route  path='/register' element = {<React.Suspense fallback={<div>loading</div>}><RegisSign/></React.Suspense> } > </Route> 
        </Routes >

      </BrowserRouter>
    );
}

export default App;