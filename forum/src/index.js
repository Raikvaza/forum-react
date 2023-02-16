import React, {useEffect, useState, createContext} from 'react';
import { BrowserRouter, Route, Routes} from "react-router-dom";
import ReactDOM from 'react-dom';
import LoginPage from './routes/Login-Page/LoginPage';
import SignUpPage from './routes/Sign-Up-Page/SignUpPage';
import HomePage from './routes/Home-Page/HomePage';
import CreatePost from './routes/Create-Post/CreatePost';
import PostPage from './routes/Post-Page/PostPage';

const AuthContext = createContext();
export default AuthContext;
function App () {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [userData, setUserData] = useState()
    const [fetchCalled, setFetchCalled] = useState(false);
      const checkAuthentication = () => {
        console.log("FETCHING");
        fetch(`http://localhost:8080/api/checkUser`, {
          headers: {
            'Accept': 'application/json',
            'Credentials': 'include'
          },
          method: "GET",
          credentials: 'include',
        }).then((r) => {
            if(r.ok){
              console.log("Token");
              setIsAuthenticated (true);
              return r.json();
            } else if (r.status === 401){
              console.log("NoToken");
              setFetchCalled(true)
              setIsAuthenticated(false)
              return;
            } else {
              setFetchCalled(true)
              throw new Error("Server error")
            }
          }).then(data => {
            if(data){setUserData(data.Username)}
            setFetchCalled(true)
          })
        }

useEffect(()=>{
    (async function() {
      await checkAuthentication();
    })();
}, [])
    
console.log("APPLICATION:"+isAuthenticated);
console.log("APPLICATION:"+userData);

if (!fetchCalled) {
    return <div>Loading...</div>;
}
console.log("FINISHED FETCH");
  return (
  <AuthContext.Provider value={{ isAuth: isAuthenticated, setIsAuth: setIsAuthenticated, username: userData}}>  
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<HomePage/>} />
        <Route path="/signup" element={<SignUpPage/>} />
        <Route path="/signin" element={<LoginPage/>} />
        {/* <Route path="/signin" element={<LoginPage isAuth={isAuthenticated}/>} /> */}

        <Route path="/post/:id" element={<PostPage />} />
        <Route path="/createpost" element={<CreatePost/>} />
      </Routes>
    </BrowserRouter>
  </AuthContext.Provider>
  )
};
const container = document.getElementById('root');
ReactDOM.render(<App />, container);