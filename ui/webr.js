
const webr = (function() {
    // function escape(str) {
    //     return str.replace(/&/g, "&amp;")
    //         .replace(/</g, "&lt;")
    //         .replace(/>/g, "&gt;")
    //         .replace(/"/g, "&quot;")
    //         .replace(/'/g, "&#039;");
    // }

    function AjaxModal(option) {
        const code="<div class='modal fade bg-light' id='webr-am' tabindex='-1' aria-labelledby='webr-aml' aria-hidden='true'>" +
            "<div class='modal-dialog'>" +
                "<div class='modal-content'>" +
                        "<div class='modal-header text-bg-secondary'>" +
                            "<h5 class='modal-title' id='webr-aml'></h5>" +
                            "<button type='button' class='btn-close' data-bs-dismiss='modal' aria-label='Close'></button>" +
                        "</div>" +
                        "<div class='modal-body' id='webr-amb'></div>" +
                    "</div>" +
                "</div>" +
            "</div>"

        if ($('#webr-am').length === 0) {
            $('body').append(code);
        }
        $('#webr-aml').text(option.title);
        const am=$('#webr-am');
        am.addClass("modal-lg");
        if (option.sm){
            am.removeClass("modal-lg");
        }

        $.ajax({
            url: option.url,
            method: option.method || 'GET',
            success: function(response) {
                $('#webr-amb').html(response);
                const modal = new bootstrap.Modal(am.get(0));
                modal.show();
            },
            error: function() {
                $('#webr-amb').html('<p>加载内容失败，请稍后再试。</p>');
                const modal = new bootstrap.Modal(am.get(0));
                modal.show();
            }
        });
    }

    function BindAjaxModal() {
        $(document).on("click",'[data-webr-ajaxmodal]',function(event){
            event.preventDefault();
            AjaxModal($(this).data('webr-ajaxmodal'))
        })
    }


    function Toast(msg){
        const code=" <div class='toast-container position-fixed top-0 start-50 translate-middle-x text-center p-3' id='webr-toast'>" +
                "<div id='webr-toast-t' class='toast' role='alert' aria-live='assertive' aria-atomic='true'>" +
                    "<div class='toast-body'></div>" +
                "</div>" +
            "</div>"

        if ($('#webr-toast').length === 0) {
            $('body').append(code);
        }
        $('#webr-toast div.toast-body').text(msg)
        const bt =new bootstrap.Toast(document.getElementById('webr-toast-t'))
        bt.show()
    }

    function Error(msg){
        const code="<div class='toast-container position-fixed start-50 translate-middle-x text-center p-3' id='webr-err' style='top:10%'>" +
                "<div id='webr-err-t' class='toast' role='alert' aria-live='assertive' aria-atomic='true'>" +
                    "<div class='toast-body' style='color:red'>" +
                     "<span class='px-3 fs-3'>😞</span>" +
                     "<span class='fs-4' id='webr-err-msg'></span>" +
                    "</div>" +
                "</div>" +
            "</div>"

        if ($('#webr-err').length === 0) {
            $('body').append(code);
        }
        $('#webr-err-msg').text(msg)
        const toast =new bootstrap.Toast(document.getElementById('webr-err-t'))
        toast.show()
    }


    function onResponse(event,resp,success,fail){
        if ( !resp.hasOwnProperty('Code') ) {
            Error("服务端响应错误：" + JSON.stringify(resp));
            return;
        }
        if (resp.Code === 0) {
            if (success){
                success(resp)
            }else{
                Toast(resp.Msg || "已提交");
                setTimeout(() => {
                    if (resp.Jump){
                        location.href = resp.Jump
                    }else if ( resp.Reload ){
                        location.reload();
                    }else if ($(event).parents("#webr-am")){
                        location.reload();
                    }
                }, 500);
            }

        } else {
            if (fail){
                fail(resp);
            }else{
                Error(resp.Msg || "保存失败")
            }

        }
    }

    const loadingSVG="<div style='text-align: center'><svg width=80 height=20>"+
    "<circle cx=10 cy=10 r=8 fill='red'>"+
       "<animate attributeName='opacity' values='1;.3;1' dur='0.9s' repeatCount='indefinite'/>"+
    "</circle>"+
    "<circle cx=40 cy=10 r=8 fill='green'>"+
        "<animate attributeName='opacity' values='1;.3;1' dur='0.9s' repeatCount='indefinite' begin='0.3s'/>"+
    "</circle>"+
    "<circle cx=70 cy=10 r=8 fill='blue'>"+
        "<animate attributeName='opacity' values='1;.3;1' dur='0.9s' repeatCount='indefinite' begin='0.6s'/>"+
    "</circle>"+
    "</svg></div>"

    function Loading(target){
        $(target).html(loadingSVG)
    }

    function JSONForm(target,onSuccess,onError) {
        $(target).on('submit', function(e) {
            e.preventDefault();
            if (typeof onSuccess == "string"){
                $(onSuccess).html(loadingSVG)
            }

            let formData = $(this).serializeJSON()
            $.ajax({
                url: $(this).attr("action"),
                method: $(this).attr("method") || "POST",
                contentType: 'application/json',
                data: JSON.stringify(formData),
                beforeSend:function(xml){
                    try{
                        xml.setRequestHeader("X-Webr",window.screen.height+"/"+window.screen.width);
                    }catch(e){}
                },
                success: function(response) {
                    if (onSuccess){
                        if (typeof onSuccess === "function"){
                            onSuccess(response);
                        }else{
                            $(onSuccess).html(response)
                        }
                    }else{
                        onResponse(e,response)
                    }
                },
                error: function(xhr, status, error) {
                    if (typeof onSuccess == "string"){
                        $(onSuccess).html("错误："+error)
                    }
                    if (onError) {
                        onError(xhr, status, error)
                    }else{
                        Error("错误："+error)
                    }
                }
            });
        });
    }

    function Confirm(title,msg,callback) {
        const code="<div class='modal fade bg-light' id='webr-cf' tabindex='-1' aria-labelledby='webr-cf-t' aria-hidden='true'>" +
                "<div class='modal-dialog'>" +
                  "<div class='modal-content'>" +
                    "<div class='modal-header'>" +
                      "<h5 class='modal-title' id='webr-cf-t'></h5>" +
                      "<button type='button' class='btn-close' data-bs-dismiss='modal' aria-label='Close'></button>" +
                    "</div>" +
                    "<div class='modal-body'></div>" +
                     "<div class='modal-footer justify-content-around'>" +
                        "<button type='button' class='btn btn-primary' id='webr-cf-ok' data-bs-dismiss='modal'>确定</button>" +
                        "<button type='button' class='btn btn-secondary' data-bs-dismiss='modal'>取消</button>" +
                    "</div>" +
                  "</div>" +
                "</div>" +
            "</div>"
        if ($('#webr-cf').length === 0) {
            $('body').append(code);
        }
        $('#webr-cf-t').text(title)
        $('#webr-cf .modal-body').html(msg)
        $('#webr-cf-ok').one("click",callback)

        const modal = new bootstrap.Modal(document.getElementById('webr-cf'));
        modal.show();
    }


    function BindAjaxConfirm(){
        $(document).on('click', 'button[data-webr-ajax]',function(event) {
            event.preventDefault();
            let option=$(this).data('webr-ajax')
            Confirm(option.title||"请确认",option.msg || "",function(e){
                e.preventDefault();
                $.ajax({
                    url: option.url,
                    method: option.method || 'POST',
                    dateType: "json",
                    success: function(response) {
                        if (option.success){
                            option.success(response)
                        }else {
                            onResponse(event,response)
                        }
                    },
                    error: function(xhr, status, error) {
                        Error("错误："+error)
                    }
                });
            })
        });
    }

    function BindInputAjaxChange(){
        $(document).on('click', 'input[data-webr-ajax]', function(event) {
            event.preventDefault();
            let option=$(this).data('webr-ajax')
            let formData=option.data || {}
            formData[$(this).attr("name")]=$(this).val()
            $.ajax({
                url: option.url,
                method: option.method || 'PUT',
                contentType: 'application/json',
                data: JSON.stringify(formData),
                success: function(response) {
                    if (option.success){
                        option.success(response)
                    }else {
                        onResponse(response)
                    }
                },
                error: function(xhr, status, error) {
                    Error("错误："+error)
                }
            });
        })
    }

    function AddRow(event,tplSec,dstSec,area){
        const obj=$(event).parents( area || "form");
        const inputs=obj.find(dstSec);
        const tplCode=$(tplSec).html();
        inputs.append(tplCode);
    }

    function RemoveRow(event,target){
       $(event).parents(target).remove()
    }

    function BindCopy(){
        if (typeof ClipboardJS === 'undefined'){
            return
        }
        const clipboard = new ClipboardJS('.copy-btn');
        const class1="text-success bi-clipboard-check-fill"
        clipboard.on('success', function(e) {
            e.clearSelection();
            Toast("已复制到剪贴板")
            const t=$(e.trigger)
            t.addClass(class1).removeClass("bi-copy")
            setTimeout(() => {
                t.addClass("bi-copy").removeClass(class1)
            }, 1500);
        });

        clipboard.on('error', function(e) {
            Error("复制失败")
        });
    }

    function BindHover(){
        $('[data-webr-hover]').on('mouseenter',function(){
            let option=$(this).data('webr-hover')
            $(this).addClass(option)
        }).on('mouseleave',function(){
            let option=$(this).data('webr-hover')
            $(this).removeClass(option)
        })
    }

    function BindSortable(){
        if (typeof Sortable === 'undefined'){
            return
        }
        const optDef={
            animation: 400,
            handle: ".sort-handle"
        }
        $('[data-webr-sortable]').each(function (){
            const value=$(this).data('webr-sortable')||{}
            const option={...optDef,...value}
            Sortable.create($(this).get(0),option);
        })
    }

    function BindCountDown(){
        $('[data-webr-countdown]').each(function (){
            const target=$(this)
            // 目标日期 如 '2025-12-31T23:59:59'
            const end = new Date(target.data('webr-countdown'));
            function update() {
                if (target.closest('body').length<1 && di){
                    clearInterval(di)
                    return
                }
                const now = new Date();
                let diff = end - now;
                let text=""
                // 如果目标日期已经过去，停止倒计时
                if (diff <= 0) {
                    text+="-"
                    diff=Math.abs(diff)
                }

                const days = Math.floor(diff / (1000 * 86400));
                const hours = Math.floor((diff % (1000 * 86400)) / (1000 * 3600));
                const minutes = Math.floor((diff % (1000 * 3600)) / (1000 * 60));
                const seconds = Math.floor((diff % (1000 * 60)) / 1000);


                if ( days>0 ){
                    text+=`${days} 天 `
                }
                if ( days>0 || hours>0 ){
                    text+=`${hours} 小时 `
                }
                if ( days>0 || hours>0 || minutes>0 ){
                    text+=`${minutes} 分钟 `
                }

                text+=`${seconds} 秒`
                target.text(text);
            }
            const di=setInterval(update, 1000);
            update();
        })

    }



    return {
        AjaxModal:AjaxModal,
        Toast:Toast,
        Error:Error,
        JSONForm:JSONForm,
        Confirm:Confirm,
        BindAjaxModal:BindAjaxModal,
        BindAjaxConfirm:BindAjaxConfirm,
        BindInputAjaxChange:BindInputAjaxChange,
        AddRow:AddRow,
        RemoveRow:RemoveRow,
        BindCopy:BindCopy,
        BindHover:BindHover,
        BindSortable:BindSortable,
        BindCountDown:BindCountDown,
        Loading:Loading,
    }
})();

jQuery().ready(function(){
    webr.BindAjaxModal()
    webr.BindAjaxConfirm()
    webr.BindInputAjaxChange()
    webr.BindCopy()
    webr.BindHover()
    webr.BindSortable()
    webr.BindCountDown()
})