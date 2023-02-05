import React, { useState } from "react";
import './Login.css'


window.onload = function() {
    (function() {
        const inputText = document.querySelectorAll('.auth-form__input');
  
        inputText.forEach( function(input) {
            input.addEventListener('focus', function() {
                this.classList.add('focus');
                this.parentElement.querySelector('.auth-form__placeholder').classList.add('focus');
            });
            input.addEventListener('blur', function() {
                this.classList.remove('focus');
                if (! this.value) {
                    this.parentElement.querySelector('.auth-form__placeholder').classList.remove('focus');
                }
            });
        });
    })();
  
    (function() {
        const togglers = document.querySelectorAll('.password-toggler');
  
        togglers.forEach( function(checkbox) {
            checkbox.addEventListener('change', function() {
  
                const toggler = this.parentElement,
                      input   = toggler.parentElement.querySelector('.input-password'),
                      icon    = toggler.querySelector('.auth-form__icon');
  
                if (checkbox.checked) {
                    input.type = 'text';
                    icon.classList.remove('la-eye')
                    icon.classList.add('la-eye-slash');
                }
  
                else
                {
                    input.type = 'password';
                    icon.classList.remove('la-eye-slash')
                    icon.classList.add('la-eye');
                }
            });
        });
    })();
  
    (function() {
        const validEmail = 'test@example.com',
              validPassword = 'qwerty123';
        
        document.body.querySelector('.hint')
                     .innerHTML = `<p>${validEmail}</p><p>${validPassword}</p>`;
  
        document.forms['form-auth'].addEventListener('submit', function(e) {
            e.preventDefault();
  
            const answerContainer = this.querySelector('.auth-form__answer'),
                  email = this.elements.email.value,
                  password = this.elements.password.value;
  
            const placeholders = document.querySelectorAll('.auth-form__placeholder');
  
            if (email === validEmail && password === validPassword) {
                answerContainer.innerHTML = '<span class="text-success">you\'ve been logged successfully</span>';
            }
  
            else {
                placeholders.forEach(function(placeholder) {
                    placeholder.classList.remove('focus');
                });
                this.elements.email.value = '';
                this.elements.password.value = '';
                answerContainer.innerHTML = '<span class="text-danger">invalid email or password</span>';
            }
        });
    })();
  };



const Login = () => {
    const [mail, setMail] = useState('');
    const [password, setPassword] = useState('');
    
    
    const sendForm = (e) => {
        e.preventDefault();
        setMail(e.target.email.value);
        setPassword(e.target.password.value);
// Check for Mail and password validity
        (async() => {
            await fetch(`http://localhost:8080/api/checkPassword`, 
            {
                headers: {
                    'Accept': 'application/json',
                    'Content-type': 'text/plain'
                },
                method: "POST",
                body: JSON.stringify({
                    Username: mail,
                    Password: password
                })
            }).then((r) => r.json())
            .then((data) => {
               console.log(data); 
            });
        })();
    }
    
    return (
        <div className="modal__background">
        <div className="modal__window">
    
            <form className="auth-form" name="form-auth" onSubmit={sendForm}>
    
                <label className="auth-form__label">
                    <span className="auth-form__placeholder">email</span>
                    <input className="auth-form__input input-email" type="email" name="email" autoComplete="off" required/>
                </label>
    
                <label className="auth-form__label">
                    <span className="auth-form__placeholder">password</span>
                    <input className="auth-form__input input-password" type="password" name="password" autocomlete="off" required/>
                    <div className="auth-form__toggler">
                        <i className="la la-eye auth-form__icon"></i>
                        <input type="checkbox" className="auth-form__checkbox password-toggler"/>
                    </div>
                </label>
    
                <div className="auth-form__answer"></div>
    
                <input className="auth-form__submit" type="submit" value="Login"/>
                
                <div className="auth-form__bottom">
                    <span>Have no account?</span>
                    <a href="#">Create new</a>
                </div>
            </form>
    
        </div>
    </div>
    
    )
}

export default Login;