/* eslint-disable jsx-a11y/anchor-is-valid */
import { useTranslation } from "react-i18next";

import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import * as Yup from 'yup'
import axios from "axios";
import RegisSign from './signup';
import { Link } from 'react-router-dom'

function RegisLog() {
  const { t } = useTranslation(["login"]);
  const formSchema = Yup.object().shape({
    email: Yup.string()
        .required('Email is mendatory')
        .matches(/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/, 'Email is invalid'),
    username: Yup.string()
    .required('username is mendatory')

        ,
    password: Yup.string()
        .required('Password is mendatory')
        .min(8, 'Password must be at 8 char long'),
    confirmPwd: Yup.string()
        .required('Password is mendatory')
        .oneOf([Yup.ref('password')], 'Passwords does not match'),

})
const formOptions = { resolver: yupResolver(formSchema) }
const { register, handleSubmit, formState } = useForm(formOptions)
const { errors } = formState
async function onSubmit(data) {
        
  console.log(data['username'])
  
  var dff= axios({url:'http://localhost:3000/ln',method:"post",data:{email:data['email'],password:data['password']},headers:{"X-API-KEY":"ax47"}});
  console.log(await dff)
  return false
}
    return (

      
      
                
                <div className="wrapper">
  <div className="title-text">
    <div className="title login">
      {t("login")}
    </div>
    
  </div>
  <div className="form-container">
    <div className="form-inner">
      <form  onSubmit={handleSubmit(onSubmit)} className="login">
        <div className="field">
          <input type="email"  name="email" {...register('email')} placeholder="Email Address"  className={`text-white ${errors.email ? 'is-invalid' : ''}`} />
          <div className="invalid-feedback red">{errors.email?.message}</div>
        </div>
        <div className="field">
          <input type="password"  name="password" {...register('password')} placeholder="Password" className={`text-white ${errors.email ? 'is-invalid' : ''}`} />
        </div>
        <div className="pass-link">
          <a href="#" style={{color: "#9FEF00"}}>Forgot password?</a>
        </div>
        <div className="field btn">
          <div className="btn-layer" />
          <input type="submit" value="Login"/>
        </div>
        <div className="signup-link">
          Not a member? <Link to="/register"><a>Signup now</a></Link>
        </div>
      </form>
      
    </div>
  </div>
</div>

          
    )
}

export default RegisLog;