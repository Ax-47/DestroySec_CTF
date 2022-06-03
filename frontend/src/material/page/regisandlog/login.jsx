import { useTranslation } from "react-i18next";
import axios from "axios";

import Cookies from 'universal-cookie';
import RegisSign from './signup';
import { Link } from 'react-router-dom'

import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import * as Yup from 'yup'





function RegisLog() {
  const cookies = new Cookies();
  const formSchema = Yup.object().shape({
    email: Yup.string()
      .required('Email is mendatory')
      .matches(/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/, 'Email is invalid'),
    password: Yup.string()
      .required('Password is mendatory')
  })
  
  const formOptions = { resolver: yupResolver(formSchema) }
      const { register, handleSubmit, formState } = useForm(formOptions)
     
      const { errors } = formState
      async function onSubmit(data) {
        
          var dff= axios({url:'http://localhost:9000/apilogin/ln',method:"post",data:{email:data['email'],password:data['password']},headers:{"X-API-KEY":"ax47"}});
          
          if (await (await dff).status ===200){
            window.location.href="/confirm"
            var $tham_po_mue_doo =await (await dff).data["jwt"];

            cookies.set('Destroy',$tham_po_mue_doo, { path: '/',SameSite:"None",secure:true });  
         
          }
        
    


          return true
      }

  const { t } = useTranslation(["login"]);
  return (


    <div className="wrapper">
      <div className="title-text">
        <div className="title login">
          {t("login")}
        </div>

      </div>
      <div className="form-container">
        <div className="form-inner">
          <form onSubmit={handleSubmit(onSubmit)} className="login">
            <div className="field mt-4">
              <input name="email" placeholder="Email Address" {...register('email')} className={`text-white ${errors.email ? 'is-invalid' : ''}`} />
              <div className="invalid-feedback red">{errors.email?.message}</div>
            </div>
            <div className="field mt-4">
              <input type="password" name="password" placeholder="Password" {...register('password')}
                className={` ${errors.password ? 'is-invalid' : ''}`} />
                <div className="invalid-feedback red">{errors.password?.message}</div>
            </div>
            <div className="pass-link mt-4">
              <a href="#" style={{ color: "#9FEF00" }}>Forgot password?</a>
            </div>
            <div className="field btn">
              <div className="btn-layer" />
              <input type="submit" value='Login' />
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