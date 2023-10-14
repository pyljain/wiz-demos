 a. 
 wiz draw -t sequencediagram -p ./handlers 
 mmdc -i ./oauth2-example/handlers/mermaid.mmd -o test.svg && open test.svg

 b.
 wiz_demos % wiz generate -a scorecard -p .
 wiz_demos % wiz generate -a docstrings --path ./oauth2-example  
 wiz review --path ./handlers 
 
wiz explain ./main.go -d 
python3 main.py 2>&1 | wiz explain

wiz chat 
- Can you please explain what does ./app.py does?
- Could you execute the code for me?
- Can you fix the issue
- Can you run it again