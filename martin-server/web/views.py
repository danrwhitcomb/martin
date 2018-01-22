''' MARTIN server web views '''
from django.http import HttpResponseNotAllowed
from django.shortcuts import redirect, render
from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required
from django.views.generic import TemplateView

@login_required()
def home(request):
    if request.method == 'GET':
        return render(request, template_name='base.html')

    return HttpResponseNotAllowed(['GET'])

def handle_login(request):
    if request.method == 'GET':
        return render(request, template_name='login.html')
    elif request.method == 'POST':
        username = request.POST['username']
        password = request.POST['password']
        user = authenticate(request, username=username, password=password)
        if user is not None:
            login(request, user)
            return redirect('/')
        else:
            # Return an 'invalid login' error message.
            return render(request, template_name='login.html',
                          context={'error': 'The username or password was incorrect'})
    else:
        return HttpResponseNotAllowed(['GET', 'POST'])

def handle_logout(request):
    logout(request)
    return redirect('/')
