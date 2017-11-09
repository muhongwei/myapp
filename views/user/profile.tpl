<header class="hero-unit">

    <div style="background-color:rgb(255, 255, 255);float:right;margin-top:0px;">
    myappversion:{{.container | GetMyappVersion}}<br>
    myappname:{{.container | GetMyappName}}
    </div>
    

    <div class="container">
        <div class="row">
            <div class="hero-text">
                <h1>Welcome to the Sitepoint / Beego App!</h1>
                <h2>This is My Test Version</h2>                
                <p>UserAgent:{{.UserAgent}}</p>
                <p>IP:{{.IP}}</p>
                <p>HostName:{{.HostName}}</p>
                <p>UserName:{{.UserName}}</p>
                <p>UserIntroduction:{{.UserIntroduction}}</p>                
            </div>
        </div>
    </div>
</header>
