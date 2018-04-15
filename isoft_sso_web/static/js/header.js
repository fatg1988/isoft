$(function () {
    if(sessionStorage.username){
        $(".login").html(sessionStorage.username);
        $(".login").mouseenter(function(){
            $(".login").html("注销");
        });
        $(".login").mouseleave(function(){
            $(".login").html(sessionStorage.username);
        });
    }

    $(".login").click(function () {
        var html = $(this).html();
        if(html == "注销"){
            window.location.href = "/user/logout";
        }else{
            window.location.href = "/user/login";
        }
    });
})