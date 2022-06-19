/* eslint-disable jsx-a11y/anchor-is-valid */

import { useTranslation } from "react-i18next";

import Cookies from 'universal-cookie';

import { Link } from 'react-router-dom'
import './sign.css'


import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import * as Yup from 'yup'
import axios from "axios";





export default function RegisSign() {
    const cookies = new Cookies();
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
        
        var dff= axios({url:'http://localhost:9000/apilogin/reg',method:"post",data:{username:data['username'],email:data['email'],password:data['password'],repassword:data['confirmPwd']},headers:{"X-API-KEY":"ax47"}});
        if (await (await dff).status ===200){
            window.location.href="/verify"
            
            var $tham_po_mue_doo =await (await dff).data["hee"];
            console.log($tham_po_mue_doo)
            cookies.set('Destroy',$tham_po_mue_doo, { path: '/',SameSite:"None",secure:true });  
         
          }
        return false
    }


    const { t } = useTranslation(["signup"]);
    return (


        <div className="wrapper">
            <div className="title-text">
                <div className="title login">
                    {t("signup")}
                </div>

            </div>
            <div className="form-container">
                <div className="form-inner">
                    <form onSubmit={handleSubmit(onSubmit)} className="login">
                    <div className="field mt-4">
                            <input placeholder="Username" name="username" {...register('username')} className={`${errors.username ? 'is-invalid' : ''}`}/>
                            <div className="invalid-feedback red">{errors.username?.message}</div>
                        </div>
                        <div className="field mt-4">
                            <input placeholder="Email Address" name="email" {...register('email')} className={`text-white ${errors.email ? 'is-invalid' : ''}`} />
                            <div className="invalid-feedback red">{errors.email?.message}</div>
                        </div>
                        
                        <div className="field mt-4">
                            <input placeholder="Password" name="password"
                                type="password"
                                {...register('password')}
                                className={`${errors.password ? 'is-invalid' : ''}`}
                            />
                            <div className="invalid-feedback red">{errors.password?.message}</div>
                        </div>
                        <div className="field mt-4">
                            <input placeholder="Confirm Password" name="confirmPwd"
                                type="password"
                                {...register('confirmPwd')}
                                className={` ${errors.confirmPwd ? 'is-invalid' : ''}`} />
                            <div className="invalid-feedback red">{errors.confirmPwd?.message}</div>
                        </div>
                        <div className="field btn">
                            <div className="btn-layer" />
                            <input type="submit" value="Sign up" />
                        </div>
                        <div className="signup-link">
                            Do you have an account? <Link to="/login"><a>Login now</a></Link>
                        </div>
                    </form>

                </div>
            </div>
        </div>


    )
};

