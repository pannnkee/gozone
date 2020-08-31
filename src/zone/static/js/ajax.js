//ajax提交(post方式提交)
function dele(url, data,success) {
    AjaxRequest(url, data,"DELETE",success);
}

//ajax提交(post方式提交)
function post(url, data,success) {
    AjaxRequest(url, data,"POST",success);
}

//ajax提交(get方式提交)
function get(url,data,success, cache) {
    AjaxRequest(url, data, "GET",success);
}

function AjaxRequest(url,data,method,success){
    $.ajax({
        url:url,
        type:method,
        data:data,
        dataType:"json",
        cache:false,
        success:success,
        error:function(data){
            Vue.$toast(data.status+"错误的请求: "+url);
        }
    })
}