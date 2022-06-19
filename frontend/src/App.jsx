import React, { Suspense } from 'react';
//import './App.css';




import { Routes, Route, BrowserRouter, Outlet } from 'react-router-dom';
import Navapp from './mainnav';



const LazyHome = React.lazy(() => import('./material/page/index/index'))
const LazyFeatures = React.lazy(() => import('./material/page/features'))
const LazyPricing = React.lazy(() => import('./material/page/pricing'))
const Lazymainnav = React.lazy(() => import('./mainnav'))

const LazyLogin = React.lazy(() => import('./material/page/regisandlog/login'))
const LazyRegis = React.lazy(() => import('./material/page/regisandlog/signup'))
const LazyConfirm = React.lazy(() => import('./material/page/confirmemail/index'))
const Lazyverify = React.lazy(() => import('./material/page/confirmemail/verify_user'))

function App() {

  return (
    <BrowserRouter >
    <div>
      <Navapp/>
      <Routes>
        
        <Route exact path='/' element={<React.Suspense fallback={<div>loading</div>}><LazyHome /></React.Suspense>} >
        </Route>
        <Route path='/features' element={<React.Suspense fallback={<div>loading</div>}><LazyFeatures /> </React.Suspense>} >
        </Route>
        <Route path='/pricing' element={<React.Suspense><LazyPricing /></React.Suspense>} >
        </Route>
        <Route path='/login' element={<React.Suspense fallback={<div>loading</div>}><LazyLogin /> </React.Suspense>} > </Route>
        <Route path='/register' element={<React.Suspense fallback={<div>loading</div>}><LazyRegis /></React.Suspense>} > </Route>
        <Route path='/confirm' element={<React.Suspense fallback={<div>loading</div>}><LazyConfirm /></React.Suspense>}></Route>
        <Route path='/verify' element={<React.Suspense fallback={<div>loading</div>}><Lazyverify /></React.Suspense>}></Route>
      </Routes >
      </div>
    </BrowserRouter>
  );
}

export default App;