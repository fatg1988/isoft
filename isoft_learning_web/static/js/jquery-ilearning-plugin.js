(function ($) {
    $.fn.showListData = function(options) {
        // 用于记录 A,B两个div的坐标位置
        var positions = {};

        var methods = {
            // 延迟效果函数
            sleep:function (numberMillis) { // numberMillis 毫秒
                var now = new Date();
                var exitTime = now.getTime() + numberMillis;
                while (true) {
                    now = new Date();
                    if (now.getTime() > exitTime)
                        return;
                }
            },
            // 判断div是否被选中
            checkChoose:function (e) {
                var inA = false;
                var inB = false;
                if(e.pageX >= positions.offsetLeftA && e.pageX <= positions.offsetLeftA + positions.offsetWidthA
                    && e.pageY >= positions.offsetTopA && e.pageY <= positions.offsetTopA + positions.offsetHeightA){
                    inA = true;
                }
                if(e.pageX >= positions.offsetLeftB && e.pageX <= positions.offsetLeftB + positions.offsetWidthB
                    && e.pageY >= positions.offsetTopB && e.pageY <= positions.offsetTopB + positions.offsetHeightB){
                    inB = true;
                }
                if(!inA && !inB){
                    // 鼠标离开事件
                    $(".showListDataRightSubArea").hide();
                    $(".showListDataLeftArea").find("li").css("background","#ffffff");
                    $(".showListDataRightArea").siblings("div").show();
                }
            },
            // 用于初始渲染样式
            renderCss:function () {
                $(".showListData").find("*").css("padding","0px");
                $(".showListDataLeftArea ul").css("margin","1px");
                $(".showListDataLeftArea a").css("line-height","28px");
                $(".showListDataLeftArea span").css({"font-size":"14px","line-height":"28px"});
                $(".showListDataRightAreaParent .showListDataRightSubArea li").css({"float":"left","display":"block","margin":"5px 5px 5px 5px"});
                $(".showListDataRightAreaParent .showListDataRightSubArea li.nofloat").css({"float":"none"});
                $(".showListDataRightAreaParent li span").css({"width":"20px","display":"inline-block","text-align":"center"});
                $(".showListDataRightAreaParent li label").css({"font-weight":"bold"});
                $(".showListDataRightAreaParent li hr").css({"margin": "0px"});
                $(".showListDataRightAreaParent .showListDataRightSubArea").css({"background-color":"#f7f7f7","display":"none","width":"100%","height":"450px",
                    "border":"2px solid #ff5000","border-left":"none"});
                $(".showListDataRightAreaParent .showListDataRightSubArea ul").css("list-style","none");
            },
            initDataEvent:function () {
                $(".showListData").find(".showListDataLeftArea").find("li").mouseover(function () {
                    $(this).css("background-color","#ffe4dc");
                    $(this).siblings().css("background","#ffffff");
                    methods.sleep(200);
                    $(".showListDataRightSubArea[control=" + $(this).attr("control") + "]").show();
                    $(".showListDataRightSubArea[control!=" + $(this).attr("control") + "]").hide();
                    $(".showListDataRightArea").siblings("div").hide();

                    positions.offsetLeftA = $(".showListDataLeftAreaParent").offset().left;
                    positions.offsetTopA = $(".showListDataLeftAreaParent").offset().top;
                    positions.offsetHeightA = $(".showListDataLeftAreaParent").height();
                    positions.offsetWidthA = $(".showListDataLeftAreaParent").width();

                    positions.offsetLeftB = $(".showListDataRightAreaParent").offset().left;
                    positions.offsetTopB = $(".showListDataRightAreaParent").offset().top;
                    positions.offsetHeightB = $(".showListDataRightAreaParent").height();
                    positions.offsetWidthB = $(".showListDataRightAreaParent").width();
                });
                $(".showListData").find(".showListDataLeftAreaParent").mouseleave(function (e) {
                    methods.checkChoose(e);
                });
                $(".showListData").find(".showListDataRightAreaParent").mouseleave(function (e) {
                    methods.checkChoose(e);
                });
                $(".showListDataRightSubArea a").click(function () {
                    var search = $(this).html();
                    window.location.href = "/course/search?search=" + search;
                });
                
            },
            initDataLeftUI:function (initDataJson) {
                var leftHtml = "";
                for(var i=0; i<initDataJson.length; i++){
                    var initDataJsonArr = initDataJson[i];
                    var left_item_arr = initDataJsonArr.left_items;
                    var appendHtml = "";
                    for(var j=0; j<left_item_arr.length; j++){
                        var left_item = left_item_arr[j];
                        appendHtml += '<a href="javascript:;">&nbsp;' + left_item + '</a>';
                        if(!(j == left_item_arr.length - 1)){
                            appendHtml += '<span>&nbsp;/&nbsp;</span>';
                        }
                    }
                    leftHtml += '<li control="index_' + i + '">' + appendHtml + '</li>';
                }
                $(".showListDataLeftArea").append("<ul>" + leftHtml + "</ul>");
            },
            initDataRightUI:function (initDataJson) {
                for(var i=0; i<initDataJson.length; i++){
                    var initDataJsonArr = initDataJson[i];
                    // 获取右边 div 所需要显示的内容
                    var right_item_arr = initDataJsonArr.right_items;
                    var appendHtml = "";
                    // 右边 div 分为 j 块
                    for(var j=0; j<right_item_arr.length; j++){
                        var right_item = right_item_arr[j];
                        // 每块有个 label
                        var right_item_label = right_item.label;
                        appendHtml += "<li class='nofloat'><ul style='margin: 0px;'><label>" + right_item_label + "</label><hr/>";
                        // 每块有多个子项
                        var right_item_list = right_item.right_item_list;
                        for(var k=0; k<right_item_list.length; k++){
                            appendHtml += '<li><a href="javascript:;">' + right_item_list[k] + '</a><span>|</span></li>';
                        }
                        appendHtml += "<div style='clear: both;'></div></li></ul>";
                    }
                    var html = '<div class="showListDataRightSubArea" control="index_' + i + '""><ul>' + appendHtml + '</ul></div>';
                    $(".showListDataRightArea").append(html);
                }
            },
            initDataUI:function (url) {         // /static/json/course.json
                $.ajax({
                    type: "GET",
                    url: url,
                    success: function(data){
                        if(data.status == "SUCCESS"){
                            var configuration_value = JSON.parse(data.configuration.configuration_value);
                            methods.initDataLeftUI(configuration_value);
                            methods.initDataRightUI(configuration_value);
                            methods.initDataEvent();
                            methods.renderCss();
                        }
                    }
                });
            }
        };

        var defaults = {};
        options = $.extend(defaults, options);
        methods.initDataUI(options.url);

        return this;
    };
})(jQuery);

$(function () {
    $(".showListData").showListData({"url":"/common/query_configuration?configuration_name=course_type"});
});

