const config = {
    WS_URL: "ws://localhost:4321/",
};

if (process.env){
    if (process.env.REACT_APP_WS_URL)  {
        config.WS_URL = process.env.REACT_APP_WS_URL;
    }
}

export default config;
