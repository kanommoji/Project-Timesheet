console.log(document.cookie);

function showSummary(){
    var memberIDByCookie = getCookie("member_id")
    var date = $("#date_summary").val(); 
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;

    const monthNames = ["JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE","JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"];
      $(document).click(function(){
        $("#title_timesheet").text(month+"-"+monthNames[month-1]+year+"-TIMESHEET");  
    });
    
    var request = new XMLHttpRequest();
    var url = "/showSummaryTimesheet";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) { 
            var json = JSON.parse(request.responseText);
            var siamChamnankit = "";
            var countSiamChamnankit = 0 ;
            var totalCoachingSiamChamnankit = 0;
            var totalTraningSiamChamnankit = 0;
            var totalOtherSiamChamnankit = 0;
            var totalIncomesSiamChamnankit = 0;
            var totalSalarySiamChamnankit = 0;
            var totalIncomeTax1SiamChamnankit = 0;
            var totalSocialSecuritySiamChamnankit = 0;
            var totalNetSalarySiamChamnankit = 0;
            var totalWageSiamChamnankit = 0;
            var totalIncomeTax53SiamChamnankit = 0;
            var totalNetWageSiamChamnankit = 0;
            var totalNetTransferSiamChamnankit = 0;
            
            var shuhari = "";
            var countShuhari = 0;
            var totalCoachingShuhari = 0;
            var totalTraningShuhari = 0;
            var totalOtherShuhari = 0;
            var totalIncomesShuhari = 0;
            var totalSalaryShuhari = 0;
            var totalIncomeTax1Shuhari = 0;
            var totalSocialSecurityShuhari = 0;
            var totalNetSalaryShuhari = 0;
            var totalWageShuhari = 0;
            var totalIncomeTax53Shuhari = 0;
            var totalNetWageShuhari = 0;
            var totalNetTransferShuhari = 0;
   
            if (json !== null){
            for (var i = 1; i <= json.length; i++) {  
                if(json[i-1].company === "siam_chamnankit"){
                    countSiamChamnankit++;
                    siamChamnankit += "<tr id=\"row_summary_id_"+i+"\">";
                    siamChamnankit += "<td>"+countSiamChamnankit+"</td>";
                    siamChamnankit += "<td id=\"member_name_th_id_"+i+"\">"+json[i-1].member_name_th+"</td>";
                    siamChamnankit += "<td id=\"coaching_id_"+i+"\">"+json[i-1].coaching.toFixed(2)+"</td>";
                    totalCoachingSiamChamnankit += json[i-1].coaching;
                    siamChamnankit += "<td id=\"training_id_"+i+"\">"+json[i-1].training.toFixed(2)+"</td>";
                    totalTraningSiamChamnankit +=json[i-1].training;
                    siamChamnankit += "<td id=\"other_id_"+i+"\">"+json[i-1].other.toFixed(2)+"</td>";
                    totalOtherSiamChamnankit += json[i-1].other;
                    siamChamnankit += "<td id=\"total_incomes_id_"+i+"\">"+json[i-1].total_incomes.toFixed(2)+"</td>";
                    totalIncomesSiamChamnankit += json[i-1].total_incomes;
                    siamChamnankit += "<td id=\"salary_id_"+i+"\">"+json[i-1].salary.toFixed(2)+"</td>";
                    totalSalarySiamChamnankit += json[i-1].salary;
                    siamChamnankit += "<td id=\"income_tax_1_id_"+i+"\">"+json[i-1].income_tax_1.toFixed(2)+"</td>";
                    totalIncomeTax1SiamChamnankit += json[i-1].income_tax_1;
                    siamChamnankit += "<td id=\"social_security_id_"+i+"\">"+json[i-1].social_security.toFixed(2)+"</td>";
                    totalSocialSecuritySiamChamnankit += json[i-1].social_security;
                    siamChamnankit += "<td id=\"net_salary_id_"+i+"\">"+json[i-1].net_salary.toFixed(2)+"</td>";
                    totalNetSalarySiamChamnankit += json[i-1].net_salary;
                    siamChamnankit += "<td id=\"wage_id_"+i+"\">"+json[i-1].wage.toFixed(2)+"</td>";
                    totalWageSiamChamnankit += json[i-1].wage;
                    siamChamnankit += "<td id=\"income_tax_53_percentage_id_"+i+"\">"+json[i-1].income_tax_53_percentage+"</td>";
                    siamChamnankit += "<td id=\"income_tax_53_id_"+i+"\">"+json[i-1].income_tax_53.toFixed(2)+"</td>";
                    totalIncomeTax53SiamChamnankit += json[i-1].income_tax_53;
                    siamChamnankit += "<td id=\"net_wage_id_"+i+"\">"+json[i-1].net_wage.toFixed(2)+"</td>";
                    totalNetWageSiamChamnankit +=json[i-1].net_wage;
                    siamChamnankit += "<td id=\"net_transfer_id_"+i+"\">"+json[i-1].net_transfer.toFixed(2)+"</td>";
                    totalNetTransferSiamChamnankit += json[i-1].net_transfer;
                    if (json[i-1].member_id == memberIDByCookie || memberIDByCookie == "001") {
                        siamChamnankit += "<td><select id=\"status_checking_transfer_"+i+"\"><option value=\""+json[i-1].status_checking_transfer+"\">"+json[i-1].status_checking_transfer+"</option>";
                        siamChamnankit += "<option value=\"รอการตรวจสอบ\">รอการตรวจสอบ</option>";
                        siamChamnankit += "<option value=\"โอนเงินเรียบร้อย\">โอนเงินเรียบร้อย</option>";
                        siamChamnankit += "<option value=\"ถูกต้อง\">ถูกต้อง</option>";
                        siamChamnankit += "<option value=\"ไม่ถูกต้อง\">ไม่ถูกต้อง</option>";
                        siamChamnankit += "</select></td>";
                        siamChamnankit += "<td><input type=\"text\" id=\"date_transfer_"+i+"\" value=\""+json[i-1].date_transfer+"\"></td>";
                        siamChamnankit += "<td><input type=\"text\" id=\"comment_"+i+"\" value=\""+json[i-1].comment+"\"></td>";
                        siamChamnankit += "<input type=\"hidden\" id=\"transaction_id_"+i+"\" value=\""+json[i-1].id+"\">";
                        siamChamnankit += "<input type=\"hidden\" id=\"member_id_"+i+"\" value=\""+json[i-1].member_id+"\">";
                        siamChamnankit += "<td>"+"<input class=\"button\" type=\"submit\" id=\"button_change_status_checking_transfer_id_"+i+"\" value=\"ยืนยันสถานะ\" onclick=\"updateStatusTransfer("+i+")\"/>"+"</td>";
                    }else{
                        siamChamnankit += "<td id=\"status_checking_transfer_"+i+"\">"+json[i-1].status_checking_transfer+"</td>";
                        siamChamnankit += "<td id=\"date_transfer_"+i+"\">"+json[i-1].date_transfer+"</td>";
                        siamChamnankit += "<td id=\"comment_"+i+"\">"+json[i-1].comment+"</td>";
                    }
                    siamChamnankit += "</tr>";
                }else{
                    countShuhari++;
                    shuhari += "<tr id=\"row_summary_id_"+i+"\">";
                    shuhari += "<td>"+countShuhari+"</td>";
                    shuhari += "<td id=\"member_name_th_id_"+i+"\">"+json[i-1].member_name_th+"</td>";
                    shuhari += "<td id=\"coaching_id_"+i+"\">"+json[i-1].coaching.toFixed(2)+"</td>";
                    totalCoachingShuhari += json[i-1].coaching;
                    shuhari += "<td id=\"training_id_"+i+"\">"+json[i-1].training.toFixed(2)+"</td>";
                    totalTraningShuhari += json[i-1].training;
                    shuhari += "<td id=\"other_id_"+i+"\">"+json[i-1].other.toFixed(2)+"</td>";
                    totalOtherShuhari += json[i-1].other;
                    shuhari += "<td id=\"total_incomes_id_"+i+"\">"+json[i-1].total_incomes.toFixed(2)+"</td>";
                    totalIncomesShuhari += json[i-1].total_incomes;
                    shuhari += "<td id=\"salary_id_"+i+"\">"+json[i-1].salary.toFixed(2)+"</td>";
                    totalSalaryShuhari += json[i-1].salary;
                    shuhari += "<td id=\"income_tax_1_id_"+i+"\">"+json[i-1].income_tax_1.toFixed(2)+"</td>";
                    totalIncomeTax1Shuhari += json[i-1].income_tax_1;
                    shuhari += "<td id=\"social_security_id_"+i+"\">"+json[i-1].social_security.toFixed(2)+"</td>";
                    totalSocialSecurityShuhari += json[i-1].social_security;
                    shuhari += "<td id=\"net_salary_id_"+i+"\">"+json[i-1].net_salary.toFixed(2)+"</td>";
                    totalNetSalaryShuhari += json[i-1].net_salary;
                    shuhari += "<td id=\"wage_id_"+i+"\">"+json[i-1].wage.toFixed(2)+"</td>";
                    totalWageShuhari += json[i-1].wage;
                    shuhari += "<td id=\"income_tax_53_percentage_id_"+i+"\">"+json[i-1].income_tax_53_percentage+"</td>";
                    shuhari += "<td id=\"income_tax_53_id_"+i+"\">"+json[i-1].income_tax_53.toFixed(2)+"</td>";
                    totalIncomeTax53Shuhari += json[i-1].income_tax_53;
                    shuhari += "<td id=\"net_wage_id_"+i+"\">"+json[i-1].net_wage.toFixed(2)+"</td>";
                    totalNetWageShuhari += json[i-1].net_wage;
                    shuhari += "<td id=\"net_transfer_id_"+i+"\">"+json[i-1].net_transfer.toFixed(2)+"</td>";
                    totalNetTransferShuhari += json[i-1].net_transfer;
                    if (json[i-1].member_id == memberIDByCookie || memberIDByCookie == "001") {
                        shuhari += "<td><select id=\"status_checking_transfer_"+i+"\"><option value=\""+json[i-1].status_checking_transfer+"\">"+json[i-1].status_checking_transfer+"</option>";
                        shuhari += "<option value=\"รอการตรวจสอบ\">รอการตรวจสอบ</option>";
                        shuhari += "<option value=\"โอนเงินเรียบร้อย\">โอนเงินเรียบร้อย</option>";
                        shuhari += "<option value=\"ถูกต้อง\">ถูกต้อง</option>";
                        shuhari += "<option value=\"ไม่ถูกต้อง\">ไม่ถูกต้อง</option>";
                        shuhari += "</select></td>";
                        shuhari += "<td><input type=\"text\" id=\"date_transfer_"+i+"\" value=\""+json[i-1].date_transfer+"\"></td>";
                        shuhari += "<td><input type=\"text\" id=\"comment_"+i+"\" value=\""+json[i-1].comment+"\"></td>";
                        shuhari += "<input type=\"hidden\" id=\"transaction_id_"+i+"\" value=\""+json[i-1].id+"\">";
                        shuhari += "<input type=\"hidden\" id=\"member_id_"+i+"\" value=\""+json[i-1].member_id+"\">";
                        shuhari += "<td>"+"<input class=\"button\" type=\"submit\" id=\"button_change_status_checking_transfer_id_"+i+"\" value=\"ยืนยันสถานะ\" onclick=\"updateStatusTransfer("+i+")\"/>"+"</td>";
                    }else{
                        shuhari += "<td id=\"status_checking_transfer_"+i+"\">"+json[i-1].status_checking_transfer+"</td>";
                        shuhari += "<td id=\"date_transfer_"+i+"\">"+json[i-1].date_transfer+"</td>";
                        shuhari += "<td id=\"comment_"+i+"\">"+json[i-1].comment+"</td>";
                    }
                    shuhari += "</tr>";
                }
            }
            $("#table_siam_chamnankit").html(siamChamnankit);
            $("#total_coaching_siam_chamnankit").html("฿ "+totalCoachingSiamChamnankit.toFixed(2));
            $("#total_traning_siam_chamnankit").html("฿ "+totalTraningSiamChamnankit.toFixed(2));
            $("#total_other_siam_chamnankit").html("฿ "+totalOtherSiamChamnankit.toFixed(2));
            $("#total_incomes_siam_chamnankit").html("฿ "+totalIncomesSiamChamnankit.toFixed(2));
            $("#total_salary_siam_chamnankit").html("฿ "+totalSalarySiamChamnankit.toFixed(2));
            $("#total_income_tax_1_siam_chamnankit").html("฿ "+totalIncomeTax1SiamChamnankit.toFixed(2));
            $("#total_social_security_siam_chamnankit").html("฿ "+totalSocialSecuritySiamChamnankit.toFixed(2));
            $("#total_net_salary_siam_chamnankit").html("฿ "+totalNetSalarySiamChamnankit.toFixed(2));
            $("#total_wage_siam_chamnankit").html("฿ "+totalWageSiamChamnankit.toFixed(2));
            $("#total_income_tax_53_siam_chamnankit").html("฿ "+totalIncomeTax53SiamChamnankit.toFixed(2));
            $("#total_net_wage_siam_chamnankit").html("฿ "+totalNetWageSiamChamnankit.toFixed(2));
            $("#total_net_transfer_siam_chamnankit").html("฿ "+totalNetTransferSiamChamnankit.toFixed(2));
            
            $("#table_shuhari").html(shuhari);
            $("#total_coaching_shuhari").html("฿ "+totalCoachingShuhari.toFixed(2));
            $("#total_traning_shuhari").html("฿ "+totalTraningShuhari.toFixed(2));
            $("#total_other_shuhari").html("฿ "+totalOtherShuhari.toFixed(2));
            $("#total_incomes_shuhari").html("฿ "+totalIncomesShuhari.toFixed(2));
            $("#total_salary_shuhari").html("฿ "+totalSalaryShuhari.toFixed(2));
            $("#total_income_tax_1_shuhari").html("฿ "+totalIncomeTax1Shuhari.toFixed(2));
            $("#total_social_security_shuhari").html("฿ "+totalSocialSecurityShuhari.toFixed(2));
            $("#total_net_salary_shuhari").html("฿ "+totalNetSalaryShuhari.toFixed(2));
            $("#total_wage_shuhari").html("฿ "+totalWageShuhari.toFixed(2));
            $("#total_income_tax_53_shuhari").html("฿ "+totalIncomeTax53Shuhari.toFixed(2));
            $("#total_net_wage_shuhari").html("฿ "+totalNetWageShuhari.toFixed(2));
            $("#total_net_transferShuhari").html("฿ "+totalNetTransferShuhari.toFixed(2));
        }  
        }
    }; 

    window.onscroll = function (){
        if (document.body.scrollTop > 10 || document.documentElement.scrollTop > 10) {
            document.getElementById("button_to_top").style.display="block";
        } else {
            document.getElementById("button_to_top").style.display="none";
        }
    };

    var data = JSON.stringify({"year":year, "month": month});
    request.send(data);
}

function updateStatusTransfer(index){
    var transactionID = $("#transaction_id_"+index).val();
    var statusTransfer = $("#status_checking_transfer_"+index).val();
    var dateTransfer = $("#date_transfer_"+index).val();
    var memberID = $("#member_id_"+index).val();
    var comment = $("#comment_"+index).val();
    var request = new XMLHttpRequest();
    var url = "/updateStatusCheckingTransfer";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token"));
    request.onreadystatechange = function () {
        if (request.status === 401){
            alert("กรุณาเข้าสู่ระบบใหม่")
            logout();
            if (deleteOauthState()){
                document.cookie = "member_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;";
                window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en")
            };
        }
    }
    var data = JSON.stringify({"member_id":memberID,"transaction_id":transactionID,"status":statusTransfer,"date":dateTransfer,"comment":comment});
    request.send(data); 
    window.location.replace(window.location.href)    
}

function addIncomeItem(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    memberID = params.get("id");

    var full = $("#day").val()
    var year = parseInt(full.split("-")[0]);
    var month = parseInt(full.split("-")[1]);
    var day = parseInt(full.split("-")[2]);
    
    var fullStartTimeAm = $("#start_time_am").val();
    var startTimeAm = new Date("January 02, 2006 "+fullStartTimeAm+" UTC");
    
    var fullEndTimeAm = $("#end_time_am").val();
    var endTimeAm = new Date("January 02, 2006 "+fullEndTimeAm+" UTC");

    var fullStartTimePm = $("#start_time_pm").val();
    var startTimePm = new Date("January 02, 2006 "+fullStartTimePm+" UTC");

    var fullEndTimePm = $("#end_time_pm").val();
    var endTimePm = new Date("January 02, 2006 "+fullEndTimePm+" UTC");

    var coachingCustomerCharging = parseFloat($("#coaching_customer_charging").val());

    var coachingPaymentRate = parseFloat($("#coaching_payment_rate").val());
    
    var trainingWage = parseFloat($("#training_wage").val());

    var otherWage = parseFloat($("#other_wage").val());

    var company = $("#company").val();

    var description = $("#description").val();
    
    var request = new XMLHttpRequest();
    var url = "/addIncomeItem";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token")); 
    request.onreadystatechange = function () {
        if (request.status === 401){
            alert("กรุณาเข้าสู่ระบบใหม่")
            logout();
            if (deleteOauthState()){
                document.cookie = "member_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;";
                window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en");
            };
        }
    }
    var data = JSON.stringify({"year":year,"month":month,"member_id":memberID,"incomes":{"day":day,"start_time_am":startTimeAm,"end_time_am":endTimeAm,"start_time_pm":startTimePm,"end_time_pm":endTimePm,"coaching_customer_charging":coachingCustomerCharging,"coaching_payment_rate":coachingPaymentRate,"training_wage":trainingWage,"other_wage":otherWage,"company":company,"description":description}});
    request.send(data); 
    window.location.replace(window.location.href); 
}

function calculatePayment() {
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    memberID = params.get("id");
    date = params.get("date");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;

    var request = new XMLHttpRequest();
    var url = "/calculatePayment";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token")); 
    request.onreadystatechange = function () {
        if (request.status === 401){
            alert("กรุณาเข้าสู่ระบบใหม่")
            logout();
            if (deleteOauthState()){
                document.cookie = "member_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;";
                window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en")
            };
        }
    }
    var data = JSON.stringify({"member_id":memberID,"year":year,"month":month});
    request.send(data); 
    window.location.replace(window.location.href); 
}

function showSummaryByID() {
    setCurrentDate();
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    memberID = params.get("id");
    date = params.get("date");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;
    var firstDay = new Date(fullDate.getFullYear(), fullDate.getMonth(), 1);
    var lastDay = new Date(fullDate.getFullYear(), fullDate.getMonth() + 1, 0);

    var memberIDByCookie = getCookie("member_id")

    var src = "https://calendar.google.com/calendar/embed?height=600&amp;wkst=1&amp;bgcolor=%23ffffff&amp;ctz=Asia%2FBangkok&amp;src=ZW4udGgjaG9saWRheUBncm91cC52LmNhbGVuZGFyLmdvb2dsZS5jb20&amp;src=dGgudGgjaG9saWRheUBncm91cC52LmNhbGVuZGFyLmdvb2dsZS5jb20&amp;color=%230B8043&amp;color=%230B8043&amp;showTz=0&amp;showPrint=0&amp;showCalendars=0&amp;showTabs=0&amp;showNav=0&amp;dates=";
    var startDate = year.toString()+month.toString()+("0" + firstDay.getDate()).slice(-2)
    var endDate = year.toString()+month.toString()+("0" + lastDay.getDate()).slice(-2)
    
    var googleCalendarURL = "<iframe src=\""+src+startDate+"/"+endDate+"\" style=\"border-width:0\" width=\"600\" height=\"400\" frameborder=\"0\" scrolling=\"no\"></iframe>";
    
    if (date == null || memberID == null) {
        alert("โปรดกรอกข้อมูลให้ครบถ้วน");
        location.href = document.referrer
    }

    const monthNamesCapital = ["JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE","JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"];
    $(document).ready(function(){
      $("#title_timesheet_by_id").text(month+"-"+monthNamesCapital[month-1]+year+"-TIMESHEET");  
    });

    
    var request = new XMLHttpRequest();
    var url = "/showTimesheetByID";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var memberNameENG = json.member_name_eng;
		    var email = json.email;
		    var totalHours = json.total_hours;
		    var totalCoachingCustomerCharging = json.total_coaching_customer_charging;
            var totalCoachingPaymentRate = json.total_coaching_payment_rate;
            var totalTrainigWage = json.total_training_wage;
            var totalOtherWage = json.total_other_wage;
            var paymentWage = json.payment_wage;
            var incomeList = "";
            
            if (json.incomes !== null) {
                for (var i = 0; i < json.incomes.length; i++) {
                    incomeList += "<tr id=\"income_company_"+json.incomes[i].company+"\">";
                    incomeList += "<td>"+json.incomes[i].day+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].start_time_am)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].end_time_am)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].start_time_pm)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].end_time_pm)+"</td>";
                    incomeList += "<td>"+json.incomes[i].overtime+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].total_hours)+"</td>";
                    incomeList += "<td>"+json.incomes[i].coaching_customer_charging.toFixed(2)+"</td>";
                    incomeList += "<td>"+json.incomes[i].coaching_payment_rate.toFixed(2)+"</td>";
                    incomeList += "<td>"+json.incomes[i].training_wage.toFixed(2)+"</td>";
                    incomeList += "<td>"+json.incomes[i].other_wage.toFixed(2)+"</td>";
                    incomeList += "<td>"+json.incomes[i].description+"</td>";
                    incomeList += "<td><input type=\"hidden\" id=\"income_id_"+i+"\" value=\""+json.incomes[i].id+"\">"
                    incomeList += "<input type=\"hidden\" id=\"member_id_"+i+"\" value=\""+json.incomes[i].member_id+"\">"
                    if (json.incomes[i].member_id == memberIDByCookie) {
                        incomeList += "<input id=\"button_delete\" type=\"submit\" value=\"ลบ\" onclick=\"deleteIncome("+i+")\"/>"+"</td>"; 
                        setDateInIncomeFormat(json.incomes[i].day+1)
                    }
                    incomeList += "</tr>";
                }
                $("#table_timesheet").html(incomeList);
            }

            
            $("#member_name_eng").html(memberNameENG);
            $("#email").html(email);
            $("#thours").html(totalHours);
            $("#total_coaching_customer_charging").html(totalCoachingCustomerCharging.toFixed(2));
            $("#total_coaching_payment_rate").html(totalCoachingPaymentRate.toFixed(2));
            $("#total_trainig_wage").html(totalTrainigWage.toFixed(2)); 
            $("#total_other_wage").html(totalOtherWage.toFixed(2)); 
            $("#payment_wage").html(paymentWage.toFixed(2));             
            if(memberID == memberIDByCookie){
                $("#th_button_calculate").html("<input class=\"button\" type=\"button\" id=\"button_calculate_payment\" value=\"คำนวณ\" onclick=\"calculatePayment()\"/>"); 
                $("#google_calendar").html(googleCalendarURL); 
            }
            
        }
    }

    window.onscroll = function (){
        if (document.body.scrollTop > 10 || document.documentElement.scrollTop > 10) {
            document.getElementById("button_to_top").style.display="block";
        } else {
            document.getElementById("button_to_top").style.display="none";
        }
    };
    
    var data = JSON.stringify({"member_id":memberID,"year":year,"month":month});
    request.send(data); 

}

function setDateInIncomeFormat(lastIndex) {
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    date = params.get("date");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;
    if (lastIndex<10){
        $("#day").val(year+"-"+month+"-0"+lastIndex); 
    }else if (lastIndex>=10&&lastIndex<=31){
        $("#day").val(year+"-"+month+"-"+lastIndex); 
    }else{
        $("#day").val(year+"-"+month+"-01"); 
    }

}

function setTableBodyAddIncomeItem(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    memberID = params.get("id");
    var memberIDByCookie = getCookie("member_id")
    var tableBody = `<tr><th>Date</th><td><input type="date" id="day"></td></tr>
    <tr><th>Start Time</th><td><input type="time" step="1" id="start_time_am" value=09:00:00 placeholder="Start Time AM"></td></tr>
    <tr><th>End Time</th><td><input type="time"  step="1" id="end_time_am" value=12:00:00 placeholder="End Time AM"></td></tr>
    <tr><th>Start Time</th><td><input type="time" step="1" id="start_time_pm" value=13:00:00 placeholder="Start Time PM"></td></tr>
    <tr><th>End Time</th><td><input type="time"  step="1" id="end_time_pm" value=18:00:00 placeholder="End Time PM"></td></tr>
    <tr><th>Coaching Customer Charging (THB)</th><td><select id="coaching_customer_charging" value=0.00>
        <option value=0.00>฿ 0.00</option>
        <option value=5000.00>฿ 5,000.00</option>
        <option value=7500.00>฿ 7,500.00</option>
        <option value=10000.00>฿ 10,000.00</option>
        <option value=15000.00>฿ 15,000.00</option>
        </select></td></tr>
    <tr><th>Coaching Payment Rate (THB)</th><td><select id="coaching_payment_rate" value=0.00>
        <option value=0.00>฿ 0.00</option>
        <option value=5000.00>฿ 5,000.00</option>
        <option value=7500.00>฿ 7,500.00</option>
        <option value=10000.00>฿ 10,000.00</option>
        <option value=15000.00>฿ 15,000.00</option>
    </select></td></tr>
    <tr><th>Training Wage (THB)</th><td><select id="training_wage" value=0.00>
        <option value=0.00>฿ 0.00</option>
        <option value=1000.00>฿ 1,000.00</option>
        <option value=2000.00>฿ 2,000.00</option>
        <option value=3000.00>฿ 3,000.00</option>
        <option value=5000.00>฿ 5,000.00</option>
        <option value=10000.00>฿ 10,000.00</option>
    </select></td></tr>
    <tr><th>Other Wage (THB)</th><td><select id="other_wage" value=0.00>
        <option value=0.00>฿ 0.00</option>
        <option value=2000.00>฿ 2,000.00</option>
        <option value=5000.00>฿ 5,000.00</option>
        <option value=7500.00>฿ 7,500.00</option>
        <option value=10000.00>฿ 10,000.00</option>
    </select></td></tr>
    <tr><th>Company</th><td><select id="company">
        <option value="siam_chamnankit">Siam Chamnankit</option>
        <option value="shuhari">Shuhari</option>
    </select></td></tr>
    <tr><th>Description</th><td><input type="text" id="description" placeholder="Description"></td></tr>
    <tr><td colspan="2"><input class="button" type="submit" id="button_add_income_item" value="เพิ่ม" onclick="addIncomeItem()"/></td></tr>`;
    $(document).ready(function(){
        if(memberIDByCookie == memberID){
            $("#table_addIncomeItem").html(tableBody);  
        }
    });    
}

function convertTimestampToTime(timestamp){
    var date = new Date(timestamp);
    datetext = date.toUTCString();
    datetext = datetext.split(' ')[4];
    return datetext
}

function deleteIncome(index){
    var incomeID = parseInt($("#income_id_"+index).val());    
    var memberID = $("#member_id_"+index).val()  
    var request = new XMLHttpRequest();
    var url = "/deleteIncomeItem";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token")); 
    request.onreadystatechange = function () {
        if (request.status === 401){
            alert("กรุณาเข้าสู่ระบบใหม่")
            logout();
            if (deleteOauthState()){
                document.cookie = "member_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;";
                window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en")
            };
        }
    }
    var data = JSON.stringify({"id":incomeID,"member_id":memberID});
    
    request.send(data);
    window.location.replace(window.location.href) 
}

function getMemberByID(){
    var urlString = document.referrer;
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    var memberID =  params.get("id");

    var memberIDByCookie = getCookie("member_id")  

    var request = new XMLHttpRequest();
    var url = "/showMemberDetailsByID";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var member = "";
            for (var i = 0; i < json.length; i++) {
                member += "<table id =\"table_member_details\">"
                member += "<tr><th>บริษัท</th><td id=\"company_id_"+i+"\">"+json[i].company+"</td></tr>";
                member += "<tr><th>ชื่อ-นามสกุล (ภาษาไทย)</th><td id=\"member_name_th_id_"+i+"\">"+json[i].member_name_th+"</td></tr>";
                member += "<tr><th>ชื่อ-นามสกุล (ภาษาอังกฤษ)</th><td id=\"member_name_eng_id_"+i+"\">"+json[i].member_name_eng+"</td></tr>";
                member += "<tr><th>E-mail</th><td id=\"email_id_"+i+"\">"+json[i].email+"</td></tr>";
                if (memberIDByCookie == memberID || memberIDByCookie == "001") {
                    member += "<tr><th>Overtime Rate</th><td><input type=\"number\" id=\"overtime_rate_id_"+i+"\" value=\""+json[i].overtime_rate.toFixed(2)+"\"></td></tr>";
                    member += "<tr><th>Rate Per Day</th><td><input type=\"number\" id=\"rate_per_day_id_"+i+"\" value=\""+json[i].rate_per_day.toFixed(2)+"\"></td></tr>";
                    member += "<tr><th>Rate Per Hour</th><td><input type=\"number\" id=\"rate_per_hour_id_"+i+"\" value=\""+json[i].rate_per_hour.toFixed(2)+"\"></td></tr>";
                    member += "<tr><th>เงินเดือน</th><td><input type=\"number\" id=\"salary_id_"+i+"\" value=\""+json[i].salary.toFixed(2)+"\"></td></tr>";
                    member += "<tr><th>หัก ณ ที่จ่าย ภ.ง.ด.1</th><td><input type=\"number\" id=\"income_tax_1_id_"+i+"\" value=\""+json[i].income_tax_1.toFixed(2)+"\"></td></tr>";
                    member += "<tr><th>ประกันสังคม</th><td><input type=\"number\" id=\"social_security_id_"+i+"\" value=\""+json[i].social_security.toFixed(2)+"\"></td></tr>";
                    member += "<tr><th>หัก ณ ที่จ่าย ภ.ง.ด.53 (ร้อยละ)</th><td><input type=\"number\" id=\"income_tax_53_percentage_id_"+i+"\" value=\""+json[i].income_tax_53_percentage+"\"></td></tr>";
                    member += "<tr><th>ประเภทของรายได้</th><td><select id=\"status_id_"+i+"\">";
                    if (json[i].status == "wage"){
                        member += "<option value=\""+json[i].status+"\">ค่าจ้างรายวัน (wage)</option>";
                        member += "<option value=\"salary\">เงินเดือน (salary)</option>";
                    }else{
                        member += "<option value=\""+json[i].status+"\">เงินเดือน (salary)</option>";
                        member += "<option value=\"wage\">ค่าจ้างรายวัน (wage)</option>";
                    }
                    member += "</select></td></tr>";
                    member += "<tr><th>ค่าเดินทาง</th><td><input type=\"number\" id=\"travel_expense_id_"+i+"\" value=\""+json[i].travel_expense.toFixed(2)+"\"></td></tr>";
                    member += "<input type=\"hidden\" id=\"member_details_id_"+i+"\" value=\""+json[i].id+"\">";
                    member += "<input type=\"hidden\" id=\"member_id_"+i+"\" value=\""+memberID+"\">";
                    member += "<tr><td></td><td><input class=\"button\" type=\"submit\" id=\"button_edit_member_id_"+i+"\" value=\"ยืนยันการแก้ไขข้อมูล\" onclick=\"editMemberDetails("+i+")\"></td></tr>";                    
                }else{
                    member += "<tr><th>Overtime Rate</th><td id=\"overtime_rate_id_"+i+"\">"+json[i].overtime_rate.toFixed(2)+"</td></tr>";
                    member += "<tr><th>Rate Per Day</th><td id=\"rate_per_day_id_"+i+"\">"+json[i].rate_per_day.toFixed(2)+"</td></tr>";
                    member += "<tr><th>Rate Per Hour</th><td id=\"rate_per_hour_id_"+i+"\">"+json[i].rate_per_hour.toFixed(2)+"</td></tr>";
                    member += "<tr><th>เงินเดือน</th><td id=\"salary_id_"+i+"\">"+json[i].salary.toFixed(2)+"</td></tr>";
                    member += "<tr><th>หัก ณ ที่จ่าย ภ.ง.ด.1</th><td id=\"income_tax_1_id_"+i+"\">"+json[i].income_tax_1.toFixed(2)+"</td></tr>";
                    member += "<tr><th>ประกันสังคม</th><td id=\"social_security_id_"+i+"\">"+json[i].social_security.toFixed(2)+"</td></tr>";
                    member += "<tr><th>หัก ณ ที่จ่าย ภ.ง.ด.53 (ร้อยละ)</th><td id=\"income_tax_53_percentage_id_"+i+"\">"+json[i].income_tax_53_percentage+"</td></tr>";
                    if (json[i].status == "wage"){
                        member += "<tr><th>ประเภทของรายได้</th><td id=\"status_id_"+i+"\">ค่าจ้างรายวัน (wage)</td></tr>";
                    }else{
                        member += "<tr><th>ประเภทของรายได้</th><td id=\"status_id_"+i+"\">เงินเดือน (salary)</td></tr>";
                    }
                    member += "<tr><th>ค่าเดินทาง</th><td id=\"travel_expense_id_"+i+"\">"+json[i].travel_expense.toFixed(2)+"</td></tr>";
                }
                member += "</table>"
                if (i+1 < json.length) {
                    member += "<br><br><br>"
                }
            }
            $("#table_member_details").html(member);
           
            
        }
    }

    window.onscroll = function (){
        if (document.body.scrollTop > 10 || document.documentElement.scrollTop > 10) {
            document.getElementById("button_to_top").style.display="block";
        } else {
            document.getElementById("button_to_top").style.display="none";
        }
    };
    
    var data = JSON.stringify({"member_id":memberID});
    request.send(data);
}

function editMemberDetails(index){    
    var id = parseInt($("#member_details_id_"+index).val());
    var memberID = $("#member_id_"+index).val();
    var memberNameTH = $("#member_name_th_id_"+index).text();
    var memberNameENG = $("#member_name_eng_id_"+index).text();
    var email = $("#email_id_"+index).text();
    var overtimeRate = parseFloat($("#overtime_rate_id_"+index).val());
    var ratePerDay = parseFloat($("#rate_per_day_id_"+index).val());
    var ratePerHour = parseFloat($("#rate_per_hour_id_"+index).val());
    var salary = parseFloat($("#salary_id_"+index).val());
    var incomeTax1 = parseFloat($("#income_tax_1_id_"+index).val());
    var socialSecurity = parseFloat($("#social_security_id_"+index).val());
    var incomeTax53Percentage = parseInt($("#income_tax_53_percentage_id_"+index).val());
    var status = $("#status_id_"+index).val();
    var travelExpense = parseFloat($("#travel_expense_id_"+index).val());

    var request = new XMLHttpRequest();
    var url = "/updateMemberDetails";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token")); 
    request.onreadystatechange = function () {
        if (request.status === 401){
            alert("กรุณาเข้าสู่ระบบใหม่")
            logout();
            if (deleteOauthState()){
                document.cookie = "member_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;";
                window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en")
            };
        }
    }
    var data = JSON.stringify({"id":id,"member_id":memberID,"member_name_th":memberNameTH,"member_name_eng":memberNameENG,
        "email":email,"overtime_rate":overtimeRate,"rate_per_day":ratePerDay,"rate_per_hour":ratePerHour,
        "salary":salary,"income_tax_1":incomeTax1,"social_security":socialSecurity,"income_tax_53_percentage":incomeTax53Percentage,
        "status":status,"travel_expense":travelExpense}); 
    
    request.send(data);
}

function setCurrentDate(){
    var currentTime = new Date();
    var currentYear = String(currentTime.getFullYear());
    var currentMonth = String(currentTime.getMonth()+1);
    var today = currentYear + "-" + currentMonth;
    $(document).ready(function(){
        $("#date_summary").val(today);  
        $("#date").val(today); 
        setInitialHome();
        if (window.location.pathname === "/home/"){
            showSummary();
        }
    });

    const monthNames = ["JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE","JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"];
    $(document).ready(function(){
        $("#title_timesheet").text(currentMonth+"-"+monthNames[currentMonth-1]+currentYear+"-TIMESHEET");  
    }); 
}

function showProfile(){
    var request = new XMLHttpRequest();
    var url = "/showProfile";
    request.open("GET", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization",getCookie("id_token"));
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var picture = "<img class=\"circular--square\" src=\""+json.picture+"\">"
            setCookie("member_id",json.member_id,30)
            $("#picture_profile").html(picture);
            $("#email_profile").html(json.email);  
        }
        if (request.status === 401){
            logout();
            if (deleteOauthState()){
                document.cookie = "member_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;";
                window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en")
            };
        }
    }   
    request.send();
}

function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i = 0; i < ca.length; i++) {
      var c = ca[i];
      while (c.charAt(0) == ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
}

function setCookie(cname, cvalue, exdays) {
    var date = new Date();
    date.setTime(date.getTime() + (exdays*24*60*60*1000));
    var expires = "expires="+ date.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function setInitialHome(){    
    var loginButton = "<a id=\"button_login\" href=\"/login\"><span class=\"glyphicon glyphicon-log-in\"></span> Login</a>"
    var logoutButton = "<a id=\"button_logout\"><span class=\"glyphicon glyphicon-log-in\"></span> Logout</a>"

    $(document).ready(function(){
        if (getCookie("oauthstate") != ""){
            $("#login").html(logoutButton);
            showProfile();
            if (getCookie("id_token") == ""){
                if (deleteOauthState()){
                    document.cookie = "member_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;";
                    window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en")
                };
            }
            $("#button_logout").click(function(){
                logout();
                if (deleteOauthState()){
                    document.cookie = "member_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;";
                    window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en")
                };
            });
        }else{  
            $("#login").html(loginButton);
        } 
    });
}

function logout(){
    var request = new XMLHttpRequest();
    var url = "/logout";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token"));
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {       
        }
    }   
    request.send();
}

function deleteOauthState(){
    var request = new XMLHttpRequest();
    var url = "/deleteOauthState";
    request.open("GET", url, true);
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {    
        }
    }   
    request.send();
    return true
}

function topFunction() {
    document.body.scrollTop = 0;
    document.documentElement.scrollTop = 0;
}