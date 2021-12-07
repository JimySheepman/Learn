def add_time(t, s, k=""):
    k.lower()
    day_count = 0
    time, meridien = t.split(" ")
    if meridien == 'AM':
        hours, minutes = time.split(':')
        add_hours, add_minutes = s.split(':')
        new_hours = int(hours) + int(add_hours)
        new_minutes = int(minutes) + int(add_minutes)
        if new_minutes >= 60:
            add_count_hours = new_minutes // 60
            new_hours += add_count_hours
            new_minutes = new_minutes % 60
            if new_minutes < 10:
                new_minutes = '0'+str(new_minutes)
        if new_hours > 12:
            new_hours = new_hours % 12
            day_count = ((new_hours // 12))
        if k == "":
            if day_count == 0:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+meridien
                return final_time+' '+k.capitalize()
            elif day_count == 1:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+'PM'
                return final_time+' '+k.capitalize()
            elif day_count % 2 == 1:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+meridien
                return final_time+' '+k.capitalize()+' '+'(next day)'
            else:
                final_time = str(new_hours)+":"+str(new_minutes)+'PM'
                return final_time+' '+k.capitalize()+' '+'({} days later)'.format((day_count+1)//2)
        else:
            if day_count == 0:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+meridien
                return final_time+', '+k.capitalize()
            elif day_count == 1:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+'PM'
                return final_time+', '+k.capitalize()
            elif day_count % 2 == 1:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+meridien
                return final_time+', '+k.capitalize()+' '+'(next day)'
            else:
                final_time = str(new_hours)+":"+str(new_minutes)+'PM'
                return final_time+', '+k.capitalize()+' '+'({} days later)'.format((day_count+1)//2)
    else:
        hours, minutes = time.split(':')
        add_hours, add_minutes = s.split(':')
        new_hours = int(hours) + int(add_hours)
        new_minutes = int(minutes) + int(add_minutes)
        if new_minutes >= 60:
            add_count_hours = new_minutes // 60
            new_hours += add_count_hours
            new_minutes = new_minutes % 60
            if new_minutes < 10:
                new_minutes = '0'+str(new_minutes)
        if new_hours > 12:
            new_hours = new_hours % 12
            day_count = (new_hours // 12)
        if k == "":
            if day_count == 0:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+meridien
                return final_time+' '+k.capitalize()
            elif day_count == 1:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+'AM'
                return final_time+' '+k.capitalize()
            elif day_count % 2 == 1:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+meridien
                return final_time+' '+k.capitalize()+' '+'(next day)'
            else:
                final_time = str(new_hours)+":"+str(new_minutes)+'AM'
                return final_time+' '+k.capitalize()+' '+'({} days later)'.format((day_count+1)//2)
        else:
            if day_count == 0:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+meridien
                return final_time+', '+k.capitalize()
            elif day_count == 1:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+'AM'
                return final_time+', '+k.capitalize()
            elif day_count % 2 == 1:
                final_time = str(new_hours)+":"+str(new_minutes)+' '+meridien
                return final_time+', '+k.capitalize()+' '+'(next day)'
            else:
                final_time = str(new_hours)+":"+str(new_minutes)+'AM'
                return final_time+', '+k.capitalize()+' '+'({} days later)'.format((day_count+1)//2)