import React, {useEffect, useState, createContext, useContext} from 'react';
import { BrowserRouter, Route, Routes} from "react-router-dom";
import ReactDOM from 'react-dom';
import LoginPage from './routes/Login-Page/LoginPage';
import SignUpPage from './routes/Sign-Up-Page/SignUpPage';
import HomePage from './routes/Home-Page/HomePage';
import CreatePost from './routes/Create-Post/CreatePost';
import PostPage from './routes/Post-Page/PostPage';

// function withAuth(Component) {
//   return function WrappedComponent(props) {
//     const [isAuthenticated, setIsAuthenticated] = useState(false);
//     const [userData, setUserData] = useState()
//     const checkAuthentication = async () => {
       
//       await fetch(`http://localhost:8080/api/checkUser`, {
//         headers: {
//           'Accept': 'application/json ',
//           'Credentials': 'include'
//         },
//         method: "GET",
//         credentials: 'include',
//       }).then((r) => {
//           if(r.ok){
//             console.log("Token");
//             setIsAuthenticated (true);
//             return r.json();
//           } else if (r.status === 401){
//             console.log("NoToken");
//             setIsAuthenticated(false)
//             return;
//           } else {
//             throw new Error("Server error")
//           }
//         }).then(data => {
//           setUserData(data.Username)
//           console.log("index:"+userData);
//         })
//       }

//     useEffect(() => {
//       checkAuthentication()
//     }, [isAuthenticated, userData]);
//     return <Component {...props} isAuthenticated={isAuthenticated} user={userData}/>;    
//   };
// }


// Wrap the App component with the withAuth HOC
//const AppWithAuth = withAuth(App);

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
  // if (!fetchCalled){
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
// const [isAuth, setAuth] = useState(false);
  // const [User, setUser] = useState('a')
  return (
    //<AuthContext.Provider value={{isAuth, setAuth, User, setUser}}>
  <AuthContext.Provider value={{ isAuth: isAuthenticated, setIsAuth: setIsAuthenticated}}>  
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<HomePage isAuth={isAuthenticated} username={userData}/>} />
        <Route path="/signup" element={<SignUpPage/>} />
        <Route path="/signin" element={<LoginPage/>} />
        {/* <Route path="/signin" element={<LoginPage isAuth={isAuthenticated}/>} /> */}

        <Route path="/post/:id" element={<PostPage />} />
        <Route path="/createpost" element={<CreatePost isAuth={isAuthenticated} username={userData}/>} />
      </Routes>
    </BrowserRouter>
  </AuthContext.Provider>
  )
};
// export default App;
const container = document.getElementById('root');
ReactDOM.render(<App />, container);

//const root = ReactDOM.createRoot(document.getElementById('root'));
// root.render(
//   <React.StrictMode>
//     <App />
//     {/* <RouterProvider router={router} /> */}
//   </React.StrictMode>
// );